package dip

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/Dipper-Labs/Dipper-Protocol/types"
	"github.com/Dipper-Labs/go-sdk/client"
	gosdktypes "github.com/Dipper-Labs/go-sdk/client/types"
	"github.com/Dipper-Labs/go-sdk/util"

	"github.com/Dipper-Labs/dip-bridge/types"
)

const (
	dipManagerAddr = "dip16qe2drpsxtdgmpw0pxhte649gzezg4e5q8zzes"
	abiFilePath    = "/Users/sun/go/src/github.com/Dipper-Labs/bridge/contracts/dip_contracts/c.abi"
	funcName       = "mintToken"
)

var (
	logger = log.New(os.Stdout, "dip.", 0)
)

type dipLand struct {
	DipCli client.Client
}

type DipLand interface {
	MintDip(tokenLockedInfo *types.MainTokenLocked, hash common.Hash) (gosdktypes.BroadcastTxResult, error)
}

func NewDipLand(sdkCfgPath string) DipLand {
	cli, err := client.NewClient(sdkCfgPath)
	if err != nil {
		log.Fatal(err)
	}

	return &dipLand{DipCli: cli}
}

func (dl *dipLand) MintDip(tokenLockedInfo *types.MainTokenLocked, hash common.Hash) (gosdktypes.BroadcastTxResult, error) {
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

	return dl.DipCli.ContractCall(dipManagerAddr, payload, sdk.NewCoin(sdk.NativeTokenName, sdk.NewInt(0)), true)
}
