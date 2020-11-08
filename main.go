package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

func AbiFromJson(json string) (*abi.ABI, error) {
	abiObj, err := abi.JSON(strings.NewReader(json))
	if err != nil {
		return nil, err
	}

	return &abiObj, nil
}

func UnpackLog(out interface{}, abiObj *abi.ABI, event string, log types.Log) error {
	if len(log.Data) > 0 {
		if err := abiObj.Unpack(out, event, log.Data); err != nil {
			return err
		}
	}

	var indexed abi.Arguments
	for _, arg := range abiObj.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	return abi.ParseTopics(out, indexed, log.Topics[1:])
}

func ParseTokenLocked(abiObj *abi.ABI, logE types.Log) (tokenLockedEvent *MainTokenLocked, err error) {
	tokenLockedEvent = new(MainTokenLocked)
	err = UnpackLog(tokenLockedEvent, abiObj, "TokenLocked", logE)
	if err != nil {
		return nil, err
	}

	return
}

func main() {
	abiObj, err := AbiFromJson(MainABI)
	if err != nil {
		log.Fatal(err)
	}

	cli, err := ethclient.Dial("ws://localhost:8546")
	if err != nil {
		log.Fatal(err)
	}

	erc20Addr := common.HexToAddress(ethDipManagerAddr)

	account := common.HexToAddress("0x0dd023d5c543054c8612a2291b647c32d5714f51")
	balance, err := cli.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			erc20Addr,
		},
	}

	logs := make(chan types.Log)
	sub, err := cli.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case logE := <-logs:
			logString, _ := logE.MarshalJSON()
			fmt.Println(string(logString))

			tokenLockedEvent, err := ParseTokenLocked(abiObj, logE)
			if err != nil {
				log.Print(err)
				continue
			}

			go mintOnDip(tokenLockedEvent, logE.TxHash)
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
