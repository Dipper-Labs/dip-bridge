package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
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
