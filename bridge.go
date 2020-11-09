package main

import (
	"log"

	"github.com/Dipper-Labs/go-sdk/client"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ethLand struct {
	EthCli *ethclient.Client
}

type EthLand interface {
}

func NewEthLand(rawUrl string) EthLand {
	cli, err := ethclient.Dial(rawUrl)
	if err != nil {
		log.Fatal(err)
	}

	return ethLand{EthCli: cli}
}

type dipLand struct {
	DipCli client.Client
}

type DipLand interface {
}

func NewDipLand(sdkCfgPath string) DipLand {
	cli, err := client.NewClient(sdkCfgPath)
	if err != nil {
		log.Fatal(err)
	}

	return dipLand{DipCli: cli}
}

type Bridge struct {
	EthLand
	DipLand
}

func NewBridge(ethRawUrl string, dipSdkCfgPath string) Bridge {
	return Bridge{NewEthLand(ethRawUrl), NewDipLand(dipSdkCfgPath)}
}
