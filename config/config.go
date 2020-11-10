package config

import (
	"io/ioutil"

	oconfig "github.com/olebedev/config"
)

const (
	// cfg for dipper network
	DefautlDipSdkCfgFileAbsPath   = "/Users/sun/go/src/github.com/Dipper-Labs/dip-bridge/config/dip_sdk.yaml"
	DefaultDipChainDipManagerAddr = "dip16qe2drpsxtdgmpw0pxhte649gzezg4e5q8zzes"
	DefaultDipChainDipManagerAbi  = "/Users/sun/go/src/github.com/Dipper-Labs/dip-bridge/contracts/dip_land_contracts/dip_manager.abi"

	// cfg for ethereum
	DefaultEthChainWsEndpoint                = "ws://localhost:8546"
	DefaultEthChainDipManagerAddr            = "0x44A25c7dD6031Fa3E9A4f60b29cE8f9c27132ac8"
	DefaultEthChainDipManagerAbi             = "/Users/sun/go/src/github.com/Dipper-Labs/dip-bridge/contracts/eth_land_contracts/dip_manager.abi"
	DefaultEthChainStartBlockNumber          = int64(10000)
	DefaultEthChainConfirmBlockCount         = int64(20)
	DefaultEthChainStartBlockNumberFromRedis = false

	// cfg for bridge
	DefaultDetectIntervalInSeconde = 1

	// cfg for redis
	DefaultRedisEndpoint = "localhost:6379"
	DefaultRedisPassword = ""
)

var (
	// cfg for dipper network
	DipSdkCfgFileAbsPath   = DefautlDipSdkCfgFileAbsPath
	DipChainDipManagerAddr = DefaultDipChainDipManagerAddr
	DipChainDipManagerAbi  = DefaultDipChainDipManagerAbi

	// cfg for ethereum
	EthChainWsEndpoint                = DefaultEthChainWsEndpoint
	EthChainDipManagerAddr            = DefaultEthChainDipManagerAddr
	EthChainDipManagerAbi             = DefaultEthChainDipManagerAbi
	EthChainStartBlockNumber          = DefaultEthChainStartBlockNumber
	EthChainConfirmBlockCount         = DefaultEthChainConfirmBlockCount
	EthChainStartBlockNumberFromRedis = DefaultEthChainStartBlockNumberFromRedis

	// cfg for bridge
	DetectIntervalInSeconde = DefaultDetectIntervalInSeconde

	// cfg for redis
	RedisEndpoint = DefaultRedisEndpoint
	RedisPassword = DefaultRedisPassword
)

func Init(configFileAbsPath string) {
	data, err := ioutil.ReadFile(configFileAbsPath)
	if err != nil {
		panic(err)
	}

	cfgInfo, err := oconfig.ParseYaml(string(data))
	if err != nil {
		panic(err)
	}

	// cfg for dipper network
	DipSdkCfgFileAbsPath, err = cfgInfo.String("dipSdk.DipSdkCfgFileAbsPath")
	if err != nil {
		panic(err)
	}

	DipChainDipManagerAddr, err = cfgInfo.String("dipChain.DipManagerAddr")
	if err != nil {
		panic(err)
	}

	DipChainDipManagerAbi, err = cfgInfo.String("dipChain.AbiFileAbsPath")
	if err != nil {
		panic(err)
	}

	// cfg for ethereum
	EthChainWsEndpoint, err = cfgInfo.String("ethChain.WsEndpoint")
	if err != nil {
		panic(err)
	}

	EthChainDipManagerAddr, err = cfgInfo.String("ethChain.DipManagerAddr")
	if err != nil {
		panic(err)
	}

	EthChainDipManagerAbi, err = cfgInfo.String("ethChain.DipManagerAbi")
	if err != nil {
		panic(err)
	}

	StartBlockNumber, err := cfgInfo.Int("ethChain.StartBlockNumber")
	if err != nil {
		panic(err)
	}
	EthChainStartBlockNumber = int64(StartBlockNumber)

	ConfirmBlockCount, err := cfgInfo.Int("ethChain.ConfirmBlockCount")
	if err != nil {
		panic(err)
	}
	EthChainConfirmBlockCount = int64(ConfirmBlockCount)

	EthChainStartBlockNumberFromRedis, err = cfgInfo.Bool("ethChain.StartBlockNumberFromRedis")
	if err != nil {
		panic(err)
	}

	// cfg for bridge
	detectIntervalInSeconde, err := cfgInfo.Int("bridge.detectIntervalInSeconde")
	if err != nil {
		panic(err)
	}
	DetectIntervalInSeconde = detectIntervalInSeconde

	// cfg for redis
	RedisEndpoint, err = cfgInfo.String("redis.Endpoint")
	if err != nil {
		panic(err)
	}

	RedisPassword, err = cfgInfo.String("redis.Password")
	if err != nil {
		panic(err)
	}
}
