package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Dipper-Labs/dip-bridge/config"
	"github.com/Dipper-Labs/dip-bridge/dip"
	"github.com/Dipper-Labs/dip-bridge/eth"
	"github.com/Dipper-Labs/dip-bridge/redis"
	"github.com/Dipper-Labs/dip-bridge/util"
)

var (
	logger = log.New(os.Stdout, "bridge-", 0)
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
			logger.Fatal("do GetEthBlockCursor failed: [", err, "]")
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
		logger.Fatal("do AbiFromFile failed: [", err, "]")
	}

	fromBlock := bridge.calcFromBlock(ctx)
	ethDipManagerAddr := common.HexToAddress(config.EthChainDipManagerAddr)

	for {
		bridge.ethHeaderBlockRWLock.RLock()
		toBlock := bridge.ethHeaderBlock - config.EthChainConfirmBlockCount
		bridge.ethHeaderBlockRWLock.RUnlock()

		if toBlock <= fromBlock {
			logger.Println("no block to process, wait 1 second")
			time.Sleep(time.Second * 1)
			continue
		}

		logs, err := bridge.QueryTokenLockedLog(ctx, ethDipManagerAddr, fromBlock, toBlock)
		if err != nil {
			logger.Fatal("do QueryTokenLockedLog failed: [", err, "], fromBlock: ", fromBlock, ", toBlock: ", toBlock)
		}

		for _, logE := range logs {
			tokenLockedInfo, err := util.ParseTokenLocked(abiObj, logE)
			if err != nil {
				logJson, _ := logE.MarshalJSON()
				logger.Fatal("do ParseTokenLocked failed: [", err, "], logE: [", string(logJson), "]")
			}

			if bridge.EthTxidExist(ctx, logE.TxHash.String()) {
				logger.Println("txId:[", logE.TxHash.String(), "] already processed")
				continue
			}

			result, err := bridge.MintDip(tokenLockedInfo, logE.TxHash)
			if err != nil {
				tokenLockedInfoJson, _ := json.Marshal(tokenLockedInfo)
				logger.Fatal("do MintDip failed: [", err, "], tokenLockedInfo: [", string(tokenLockedInfoJson), "], txHash: ", logE.TxHash.String())
			}

			dipReceipt, err := json.Marshal(result)
			if err != nil {
				logger.Fatal("do Marshal failed: [", err, "], dipper network txid: ", result.CommitResult.Hash.String())
			}

			bridge.SaveEthTxidProcessReceiptOnDip(ctx, logE.TxHash.String(), string(dipReceipt))
			logger.Println("txId:[", logE.TxHash.String(), "] finished")
		}

		bridge.SetEthBlockCursor(ctx, toBlock)
		logger.Println("finished eth block: ", toBlock)
		fromBlock = toBlock + 1
	}
}
