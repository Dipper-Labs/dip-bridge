package main

import (
	"context"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	"github.com/Dipper-Labs/go-sdk/client"
	"github.com/Dipper-Labs/go-sdk/util"
)

const (
	ethDipManagerAddr = "0x44A25c7dD6031Fa3E9A4f60b29cE8f9c27132ac8"
)

var (
	erc20Addr   = common.HexToAddress(ethDipManagerAddr)
	headerBlock = int64(0) // 订阅以太坊，时时更新
	gRWLock     = new(sync.RWMutex)
)

func QueryLog(cli *ethclient.Client, ctx context.Context, fromBlock, toBlock int64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			erc20Addr,
		},
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		//Topics:
	}

	return cli.FilterLogs(ctx, query)
}

func LogProcess(cli *ethclient.Client, ctx context.Context) {
	abiObj, err := AbiFromJson(MainABI)
	if err != nil {
		log.Fatal(err)
	}

	startBlock := int64(1000) // 配置文件读取

	for {
		gRWLock.RLock()
		toBlock := headerBlock - 50 - startBlock
		gRWLock.RUnlock()
		if toBlock <= 0 {
			time.Sleep(time.Second * 5)
			continue
		}

		logs, err := QueryLog(cli, ctx, startBlock, toBlock)
		if err != nil {
			log.Print(err)
			continue
		}

		for _, logE := range logs {
			tokenLockedInfo, err := ParseTokenLocked(abiObj, logE)
			if err != nil {
				log.Fatal(err)
			}
			mintOnDip(tokenLockedInfo, logE.TxHash)
			time.Sleep(time.Second * 10)
		}
	}
}

func main() {
	cli, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	headerChan := make(chan *types.Header)
	headerSub, err := cli.SubscribeNewHead(context.Background(), headerChan)
	if err != nil {
		log.Fatal(err)
	}

	go LogProcess(cli, context.Background())

	for {
		select {
		case err := <-headerSub.Err():
			log.Fatal(err)
		case newHeader := <-headerChan:
			headerJson, err := newHeader.MarshalJSON()
			if err != nil {
				log.Print(err)
			}
			gRWLock.Lock()
			headerBlock = newHeader.Number.Int64()
			gRWLock.Unlock()
			log.Print(string(headerJson))
		}
	}
}

func mintOnDip(tokenLockedInfo *MainTokenLocked, hash common.Hash) {
	const dipSdkCfgFilePath = "/Users/sun/go/src/github.com/Dipper-Labs/bridge/config/sdk.yaml"
	const dipManagerAddr = "dip16qe2drpsxtdgmpw0pxhte649gzezg4e5q8zzes"
	cli, err := client.NewClient(dipSdkCfgFilePath)
	if err != nil {
		log.Fatal(err)
	}

	const abiFilePath = "/Users/sun/go/src/github.com/Dipper-Labs/bridge/contracts/dip_contracts/c.abi"
	const funcName = "mintToken"

	toAddr, err := sdk.AccAddressFromBech32(tokenLockedInfo.To)
	if err != nil {
		log.Fatal(err)
	}

	var p [20]byte
	copy(p[:], toAddr)
	payload, err := util.BuildPayloadByABIFile(abiFilePath, funcName, hash, p, tokenLockedInfo.Amount)
	if err != nil {
		log.Fatal(err)
	}

	result, err := cli.ContractCall(dipManagerAddr, payload, sdk.NewCoin(sdk.NativeTokenName, sdk.NewInt(0)), true)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(result)
}
