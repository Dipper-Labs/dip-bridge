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
			log.Println(fmt.Sprintf("ping %d secs", config.DetectIntervalInSeconde))
			time.Sleep(time.Second * time.Duration(config.DetectIntervalInSeconde))
			continue
		}

		logs, err := bridge.QueryTokenLockedLog(ctx, ethDipManagerAddr, fromBlock, toBlock)
		if err != nil {
			log.Fatalf("do QueryTokenLockedLog failed:[%v],fromBlock:%v, toBlock:%v\n", err, fromBlock, toBlock)
		}

		for _, logE := range logs {
			tokenLockedInfo, err := util.ParseTokenLocked(abiObj, logE)
			if err != nil {
				logJson, _ := logE.MarshalJSON()
				log.Fatalf("do ParseTokenLocked failed:[%v],logE:[%s]\n", err, string(logJson))
			}

			if bridge.EthTxidExist(ctx, logE.TxHash.String()) {
				log.Println("txId:[", logE.TxHash.String(), "] already processed")
				continue
			}

			result, err := bridge.MintDip(tokenLockedInfo, logE.TxHash)
			if err != nil {
				tokenLockedInfoJson, _ := json.Marshal(tokenLockedInfo)
				log.Fatalf("do MintDip failed:[%v],tokenLockedInfo:[%s],txHash:%s\n", err, string(tokenLockedInfoJson), logE.TxHash.String())
			}

			dipReceipt, err := json.Marshal(result)
			if err != nil {
				log.Fatalf("do Marshal failed:[%v],dipper network txid:%s\n", err, result.CommitResult.Hash.String())
			}

			bridge.SaveEthTxidProcessReceiptOnDip(ctx, logE.TxHash.String(), string(dipReceipt))
			log.Println("txId:[", logE.TxHash.String(), "] finished")
		}

		bridge.SetEthBlockCursor(ctx, toBlock)
		log.Println("finished eth block: ", toBlock)
		fromBlock = toBlock + 1
	}
}
