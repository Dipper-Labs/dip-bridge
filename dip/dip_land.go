package dip

import (
	"log"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"

	"github.com/Dipper-Labs/go-sdk/client"
	gosdktypes "github.com/Dipper-Labs/go-sdk/client/types"
	"github.com/Dipper-Labs/go-sdk/util"

	"github.com/Dipper-Labs/dip-bridge/config"
	"github.com/Dipper-Labs/dip-bridge/types"
)

const (
	funcName = "mintToken"
)

type dipLand struct {
	DipCli client.Client
}

type DipLand interface {
	MintDip(tokenLockedInfo *types.TypesTokenLocked, hash common.Hash) (gosdktypes.BroadcastTxResult, error)
}

func NewDipLand(sdkCfgPath string) DipLand {
	cli, err := client.NewClient(sdkCfgPath)
	if err != nil {
		log.Fatal("do NewClient failed: [", err, "]")
	}

	return &dipLand{DipCli: cli}
}

func (dl *dipLand) MintDip(tokenLockedInfo *types.TypesTokenLocked, hash common.Hash) (gosdktypes.BroadcastTxResult, error) {
	toAddr, err := sdk.AccAddressFromBech32(tokenLockedInfo.To)
	if err != nil {
		log.Fatal("do AccAddressFromBech32 failed: [", err, "], to: [", tokenLockedInfo.To, "]")
	}

	var p [20]byte
	copy(p[:], toAddr)
	payload, err := util.BuildPayloadByABIFile(config.DipChainDipManagerAbi, funcName, hash, p, tokenLockedInfo.Amount)
	if err != nil {
		log.Fatal("do BuildPayloadByABIFile failed: [", err, "], hash: [", hash.String(), "], to: [", tokenLockedInfo.To, "], amount: [", tokenLockedInfo.Amount.String(), "]")
	}

	return dl.DipCli.ContractCall(config.DipChainDipManagerAddr, payload, sdk.NewCoin(sdk.NativeTokenName, sdk.NewInt(0)), true)
}
