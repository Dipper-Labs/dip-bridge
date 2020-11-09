package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/Dipper-Labs/dip-bridge/types"
)

func UnpackEthLog(out interface{}, abiObj *abi.ABI, event string, log ethtypes.Log) error {
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

func ParseTokenLocked(abiObj *abi.ABI, logE ethtypes.Log) (tokenLockedEvent *types.MainTokenLocked, err error) {
	tokenLockedEvent = new(types.MainTokenLocked)
	err = UnpackEthLog(tokenLockedEvent, abiObj, "TokenLocked", logE)
	if err != nil {
		return nil, err
	}

	return
}
