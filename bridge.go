package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Dipper-Labs/dip-bridge/config"
	"github.com/Dipper-Labs/dip-bridge/dip"
	"github.com/Dipper-Labs/dip-bridge/eth"
	"github.com/Dipper-Labs/dip-bridge/redis"
	"github.com/Dipper-Labs/dip-bridge/util"
)

type Bridge struct {
	eth.EthLand
	dip.DipLand
	redis.RedisCli
	ethHeaderBlock       int64
	ethHeaderBlockRWLock *sync.RWMutex
}

func NewBridge(cfgPath string) *Bridge {
	config.Init(cfgPath)
	return &Bridge{
		eth.NewEthLand(config.EthChainWsEndpoint),
		dip.NewDipLand(config.DipSdkCfgFileAbsPath),
		redis.NewRedisCli(config.RedisEndpoint),
		0,
		new(sync.RWMutex),
	}
}

func (bridge *Bridge) UpdateEthHeaderBlock(HeaderBlock int64) {
	bridge.ethHeaderBlockRWLock.Lock()
	defer bridge.ethHeaderBlockRWLock.Unlock()
	bridge.ethHeaderBlock = HeaderBlock
}

func (bridge *Bridge) calcFromBlock(ctx context.Context) int64 {
	fromBlock := config.EthChainStartBlockNumber

	if config.EthChainStartBlockNumberFromRedis {
		ethBlockCursor, err := bridge.GetEthBlockCursor(ctx)
		if err != nil {
			log.Fatalf("do GetEthBlockCursor failed:[%v]\n", err)
		}

		if ethBlockCursor > 0 {
			fromBlock = ethBlockCursor
		}
	}

	return fromBlock
}

func (bridge *Bridge) RunBridge(ctx context.Context) {
	abiObj, err := util.AbiFromFile(config.EthChainDipManagerAbi)
	if err != nil {
		log.Fatalf("do AbiFromFile failed:[%v]\n", err)
	}

	fromBlock := bridge.calcFromBlock(ctx)
	ethDipManagerAddr := common.HexToAddress(config.EthChainDipManagerAddr)

	for {
		bridge.ethHeaderBlockRWLock.RLock()
		toBlock := bridge.ethHeaderBlock - config.EthChainConfirmBlockCount
		bridge.ethHeaderBlockRWLock.RUnlock()

		if toBlock <= fromBlock {
			log.Printf("ping %d secs", config.DetectIntervalInSeconde)
			time.Sleep(time.Second * time.Duration(config.DetectIntervalInSeconde))
			continue
		}

		logs, err := bridge.QueryTokenLockedLog(ctx, ethDipManagerAddr, fromBlock, toBlock)
		if err != nil {
			log.Printf("do QueryTokenLockedLog failed:[%v],fromBlock:%v, toBlock:%v\n", err, fromBlock, toBlock)
			time.Sleep(time.Second * time.Duration(10))
			continue
		}

		if len(logs) > 0 {
			log.Printf("got %d Token Locked Event between block[%v, %v]", len(logs), fromBlock, toBlock)
		}

		logsCount := len(logs)
		logIndex := 0
		for _, logE := range logs {
			logIndex++

			tokenLockedInfo, err := util.ParseTokenLocked(abiObj, logE)
			if err != nil {
				logJson, _ := logE.MarshalJSON()
				log.Fatalf("do ParseTokenLocked failed:[%v],logE:[%s]\n", err, string(logJson))
			}
			log.Printf("[%v/%v]-%v:%s-event[%s:%s:%v]", logIndex, logsCount, logE.BlockNumber, logE.TxHash.String(), tokenLockedInfo.From.String(), tokenLockedInfo.To, tokenLockedInfo.Amount.String())

			if bridge.EthTxidExist(ctx, logE.TxHash.String()) {
				log.Printf("txId:[%s] already processed", logE.TxHash.String())
				continue
			}

			result, err := bridge.MintDip(tokenLockedInfo, logE.TxHash)
			if err != nil {
				tokenLockedInfoJson, _ := json.Marshal(tokenLockedInfo)
				failedInfo := fmt.Sprintf("do MintDip failed:[%v],tokenLockedInfo:[%s],txHash:%s\n", err, string(tokenLockedInfoJson), logE.TxHash.String())
				log.Println(failedInfo)
				bridge.SaveEthTxidProcessReceiptOnDip(ctx, logE.TxHash.String(), "failed")
				bridge.SaveEthTxidProcessReceiptOnDip(ctx, fmt.Sprintf("failed.%s", logE.TxHash.String()), failedInfo)
				continue
			}

			dipReceipt, err := json.Marshal(result)
			if err != nil {
				log.Fatalf("do Marshal failed:[%v],dipper network txid:%s\n", err, result.CommitResult.Hash.String())
			}

			bridge.SaveEthTxidProcessReceiptOnDip(ctx, logE.TxHash.String(), string(dipReceipt))
			log.Printf("txId:[%s] finished", logE.TxHash.String())
		}

		bridge.SetEthBlockCursor(ctx, toBlock)
		log.Printf("finished eth block: %v", toBlock)
		fromBlock = toBlock + 1
	}
}
