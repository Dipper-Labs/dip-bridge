// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package types

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TypesABI is the input ABI used to generate the binding from.
const TypesABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenLocked\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"dipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"lockInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"substring\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Types is an auto generated Go binding around an Ethereum contract.
type Types struct {
	TypesCaller     // Read-only binding to the contract
	TypesTransactor // Write-only binding to the contract
	TypesFilterer   // Log filterer for contract events
}

// TypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypesSession struct {
	Contract     *Types            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypesCallerSession struct {
	Contract *TypesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypesTransactorSession struct {
	Contract     *TypesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypesRaw struct {
	Contract *Types // Generic contract binding to access the raw methods on
}

// TypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypesCallerRaw struct {
	Contract *TypesCaller // Generic read-only contract binding to access the raw methods on
}

// TypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypesTransactorRaw struct {
	Contract *TypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypes creates a new instance of Types, bound to a specific deployed contract.
func NewTypes(address common.Address, backend bind.ContractBackend) (*Types, error) {
	contract, err := bindTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Types{TypesCaller: TypesCaller{contract: contract}, TypesTransactor: TypesTransactor{contract: contract}, TypesFilterer: TypesFilterer{contract: contract}}, nil
}

// NewTypesCaller creates a new read-only instance of Types, bound to a specific deployed contract.
func NewTypesCaller(address common.Address, caller bind.ContractCaller) (*TypesCaller, error) {
	contract, err := bindTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypesCaller{contract: contract}, nil
}

// NewTypesTransactor creates a new write-only instance of Types, bound to a specific deployed contract.
func NewTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*TypesTransactor, error) {
	contract, err := bindTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypesTransactor{contract: contract}, nil
}

// NewTypesFilterer creates a new log filterer instance of Types, bound to a specific deployed contract.
func NewTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*TypesFilterer, error) {
	contract, err := bindTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypesFilterer{contract: contract}, nil
}

// bindTypes binds a generic wrapper to an already deployed contract.
func bindTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Types *TypesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Types.Contract.TypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Types *TypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Types.Contract.TypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Types *TypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Types.Contract.TypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Types *TypesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Types.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Types *TypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Types.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Types *TypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Types.Contract.contract.Transact(opts, method, params...)
}

// Erc20Addr is a free data retrieval call binding the contract method 0x317e3ce0.
//
// Solidity: function erc20Addr() constant returns(address)
func (_Types *TypesCaller) Erc20Addr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Types.contract.Call(opts, out, "erc20Addr")
	return *ret0, err
}

// Erc20Addr is a free data retrieval call binding the contract method 0x317e3ce0.
//
// Solidity: function erc20Addr() constant returns(address)
func (_Types *TypesSession) Erc20Addr() (common.Address, error) {
	return _Types.Contract.Erc20Addr(&_Types.CallOpts)
}

// Erc20Addr is a free data retrieval call binding the contract method 0x317e3ce0.
//
// Solidity: function erc20Addr() constant returns(address)
func (_Types *TypesCallerSession) Erc20Addr() (common.Address, error) {
	return _Types.Contract.Erc20Addr(&_Types.CallOpts)
}

// LockInfo is a free data retrieval call binding the contract method 0x10f3b183.
//
// Solidity: function lockInfo(address , string ) constant returns(uint256)
func (_Types *TypesCaller) LockInfo(opts *bind.CallOpts, arg0 common.Address, arg1 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Types.contract.Call(opts, out, "lockInfo", arg0, arg1)
	return *ret0, err
}

// LockInfo is a free data retrieval call binding the contract method 0x10f3b183.
//
// Solidity: function lockInfo(address , string ) constant returns(uint256)
func (_Types *TypesSession) LockInfo(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Types.Contract.LockInfo(&_Types.CallOpts, arg0, arg1)
}

// LockInfo is a free data retrieval call binding the contract method 0x10f3b183.
//
// Solidity: function lockInfo(address , string ) constant returns(uint256)
func (_Types *TypesCallerSession) LockInfo(arg0 common.Address, arg1 string) (*big.Int, error) {
	return _Types.Contract.LockInfo(&_Types.CallOpts, arg0, arg1)
}

// Substring is a free data retrieval call binding the contract method 0x1dcd9b55.
//
// Solidity: function substring(string str, uint256 startIndex, uint256 endIndex) constant returns(string)
func (_Types *TypesCaller) Substring(opts *bind.CallOpts, str string, startIndex *big.Int, endIndex *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Types.contract.Call(opts, out, "substring", str, startIndex, endIndex)
	return *ret0, err
}

// Substring is a free data retrieval call binding the contract method 0x1dcd9b55.
//
// Solidity: function substring(string str, uint256 startIndex, uint256 endIndex) constant returns(string)
func (_Types *TypesSession) Substring(str string, startIndex *big.Int, endIndex *big.Int) (string, error) {
	return _Types.Contract.Substring(&_Types.CallOpts, str, startIndex, endIndex)
}

// Substring is a free data retrieval call binding the contract method 0x1dcd9b55.
//
// Solidity: function substring(string str, uint256 startIndex, uint256 endIndex) constant returns(string)
func (_Types *TypesCallerSession) Substring(str string, startIndex *big.Int, endIndex *big.Int) (string, error) {
	return _Types.Contract.Substring(&_Types.CallOpts, str, startIndex, endIndex)
}

// LockToken is a paid mutator transaction binding the contract method 0xa44ce259.
//
// Solidity: function LockToken(string dipAddr, uint256 amount) returns()
func (_Types *TypesTransactor) LockToken(opts *bind.TransactOpts, dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Types.contract.Transact(opts, "LockToken", dipAddr, amount)
}

// LockToken is a paid mutator transaction binding the contract method 0xa44ce259.
//
// Solidity: function LockToken(string dipAddr, uint256 amount) returns()
func (_Types *TypesSession) LockToken(dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Types.Contract.LockToken(&_Types.TransactOpts, dipAddr, amount)
}

// LockToken is a paid mutator transaction binding the contract method 0xa44ce259.
//
// Solidity: function LockToken(string dipAddr, uint256 amount) returns()
func (_Types *TypesTransactorSession) LockToken(dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Types.Contract.LockToken(&_Types.TransactOpts, dipAddr, amount)
}

// TypesTokenLockedIterator is returned from FilterTokenLocked and is used to iterate over the raw logs and unpacked data for TokenLocked events raised by the Types contract.
type TypesTokenLockedIterator struct {
	Event *TypesTokenLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TypesTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TypesTokenLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TypesTokenLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TypesTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TypesTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TypesTokenLocked represents a TokenLocked event raised by the Types contract.
type TypesTokenLocked struct {
	From   common.Address
	To     string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenLocked is a free log retrieval operation binding the contract event 0x865128631edc31d5501681b824f124e2f7718e3569d0bad93872617fce9f97f8.
//
// Solidity: event TokenLocked(address indexed from, string to, uint256 amount)
func (_Types *TypesFilterer) FilterTokenLocked(opts *bind.FilterOpts, from []common.Address) (*TypesTokenLockedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Types.contract.FilterLogs(opts, "TokenLocked", fromRule)
	if err != nil {
		return nil, err
	}
	return &TypesTokenLockedIterator{contract: _Types.contract, event: "TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTokenLocked is a free log subscription operation binding the contract event 0x865128631edc31d5501681b824f124e2f7718e3569d0bad93872617fce9f97f8.
//
// Solidity: event TokenLocked(address indexed from, string to, uint256 amount)
func (_Types *TypesFilterer) WatchTokenLocked(opts *bind.WatchOpts, sink chan<- *TypesTokenLocked, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Types.contract.WatchLogs(opts, "TokenLocked", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TypesTokenLocked)
				if err := _Types.contract.UnpackLog(event, "TokenLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenLocked is a log parse operation binding the contract event 0x865128631edc31d5501681b824f124e2f7718e3569d0bad93872617fce9f97f8.
//
// Solidity: event TokenLocked(address indexed from, string to, uint256 amount)
func (_Types *TypesFilterer) ParseTokenLocked(log types.Log) (*TypesTokenLocked, error) {
	event := new(TypesTokenLocked)
	if err := _Types.contract.UnpackLog(event, "TokenLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}
