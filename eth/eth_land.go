package eth

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethLand struct {
	EthCli *ethclient.Client
}

type EthLand interface {
	SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription)
	QueryTokenLockedLog(ctx context.Context, ethManagerAddr common.Address, fromBlock, toBlock int64) ([]types.Log, error)
}

func NewEthLand(rawUrl string) EthLand {
	cli, err := ethclient.Dial(rawUrl)
	if err != nil {
		log.Fatal("do ethclient.Dial failed: [", err, "], rawUrl: [", rawUrl, "]")
	}

	return &ethLand{EthCli: cli}
}

func (el *ethLand) SubscribeNewHead(ctx context.Context) (chan *types.Header, ethereum.Subscription) {
	headerChan := make(chan *types.Header)
	headerSub, err := el.EthCli.SubscribeNewHead(ctx, headerChan)
	if err != nil {
		log.Fatal("do SubscribeNewHead failed: [", err, "]")
	}

	return headerChan, headerSub
}

func (el *ethLand) QueryTokenLockedLog(ctx context.Context, ethManagerAddr common.Address, fromBlock, toBlock int64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			ethManagerAddr,
		},
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
	}

	return el.EthCli.FilterLogs(ctx, query)
}
