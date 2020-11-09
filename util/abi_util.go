package util

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func AbiFromFile(file string) (*abi.ABI, error) {
	abiJson, err := LoadABI(file)
	if err != nil {
		return nil, err
	}
	abiObj, err := AbiFromJson(abiJson)
	if err != nil {
		return nil, err
	}
	return abiObj, nil
}

func AbiFromJson(json string) (*abi.ABI, error) {
	abiObj, err := abi.JSON(strings.NewReader(json))
	if err != nil {
		return nil, err
	}

	return &abiObj, nil
}

func LoadABI(abiFileAbsPath string) (string, error) {
	fd, err := os.Open(abiFileAbsPath)
	defer fd.Close()

	if err != nil {
		return "", nil
	}

	d, err := ioutil.ReadAll(fd)
	return string(d), nil
}
