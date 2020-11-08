// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenLocked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc20Addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"dipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"dipAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unlockToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MainBin is the compiled bytecode used for deploying new contracts.
var MainBin = "0x608060405234801561001057600080fd5b506040516109393803806109398339818101604052602081101561003357600080fd5b810190808051906020019092919050505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506108a4806100956000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806323ae0e0b14610046578063317e3ce01461010b57806333fb6c1d14610155575b600080fd5b6101096004803603604081101561005c57600080fd5b810190808035906020019064010000000081111561007957600080fd5b82018360208201111561008b57600080fd5b803590602001918460018302840111640100000000831117156100ad57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019092919050505061021a565b005b61011361021e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102186004803603604081101561016b57600080fd5b810190808035906020019064010000000081111561018857600080fd5b82018360208201111561019a57600080fd5b803590602001918460018302840111640100000000831117156101bc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610244565b005b5050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506102983330848473ffffffffffffffffffffffffffffffffffffffff166104a9909392919063ffffffff16565b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846040518082805190602001908083835b6020831061030b57805182526020820191506020810190506020830392506102e8565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020846040518082805190602001908083835b602083106103b2578051825260208201915060208101905060208303925061038f565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055503373ffffffffffffffffffffffffffffffffffffffff167f865128631edc31d5501681b824f124e2f7718e3569d0bad93872617fce9f97f884846040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561046957808201518184015260208101905061044e565b50505050905090810190601f1680156104965780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2505050565b6105a9848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506105af565b50505050565b6105ce8273ffffffffffffffffffffffffffffffffffffffff166107fa565b610640576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b6020831061068f578051825260208201915060208101905060208303925061066c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d80600081146106f1576040519150601f19603f3d011682016040523d82523d6000602084013e6106f6565b606091505b50915091508161076e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b6000815111156107f45780806020019051602081101561078d57600080fd5b81019080805190602001909291905050506107f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610846602a913960400191505060405180910390fd5b5b50505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561083c57506000801b8214155b9250505091905056fe5361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a72315820350f282b06a7f60ec8a2f2fc3997c0294a5d0bfebce91299575b9a5e0949efab64736f6c63430005110032"

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend, addr common.Address) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MainBin), backend, addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Erc20Addr is a free data retrieval call binding the contract method 0x317e3ce0.
//
// Solidity: function erc20Addr() constant returns(address)
func (_Main *MainCaller) Erc20Addr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "erc20Addr")
	return *ret0, err
}

// Erc20Addr is a free data retrieval call binding the contract method 0x317e3ce0.
//
// Solidity: function erc20Addr() constant returns(address)
func (_Main *MainSession) Erc20Addr() (common.Address, error) {
	return _Main.Contract.Erc20Addr(&_Main.CallOpts)
}

// Erc20Addr is a free data retrieval call binding the contract method 0x317e3ce0.
//
// Solidity: function erc20Addr() constant returns(address)
func (_Main *MainCallerSession) Erc20Addr() (common.Address, error) {
	return _Main.Contract.Erc20Addr(&_Main.CallOpts)
}

// LockToken is a paid mutator transaction binding the contract method 0x33fb6c1d.
//
// Solidity: function lockToken(string dipAddr, uint256 amount) returns()
func (_Main *MainTransactor) LockToken(opts *bind.TransactOpts, dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "lockToken", dipAddr, amount)
}

// LockToken is a paid mutator transaction binding the contract method 0x33fb6c1d.
//
// Solidity: function lockToken(string dipAddr, uint256 amount) returns()
func (_Main *MainSession) LockToken(dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.LockToken(&_Main.TransactOpts, dipAddr, amount)
}

// LockToken is a paid mutator transaction binding the contract method 0x33fb6c1d.
//
// Solidity: function lockToken(string dipAddr, uint256 amount) returns()
func (_Main *MainTransactorSession) LockToken(dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.LockToken(&_Main.TransactOpts, dipAddr, amount)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x23ae0e0b.
//
// Solidity: function unlockToken(string dipAddr, uint256 amount) returns()
func (_Main *MainTransactor) UnlockToken(opts *bind.TransactOpts, dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "unlockToken", dipAddr, amount)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x23ae0e0b.
//
// Solidity: function unlockToken(string dipAddr, uint256 amount) returns()
func (_Main *MainSession) UnlockToken(dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.UnlockToken(&_Main.TransactOpts, dipAddr, amount)
}

// UnlockToken is a paid mutator transaction binding the contract method 0x23ae0e0b.
//
// Solidity: function unlockToken(string dipAddr, uint256 amount) returns()
func (_Main *MainTransactorSession) UnlockToken(dipAddr string, amount *big.Int) (*types.Transaction, error) {
	return _Main.Contract.UnlockToken(&_Main.TransactOpts, dipAddr, amount)
}

// MainTokenLockedIterator is returned from FilterTokenLocked and is used to iterate over the raw logs and unpacked data for TokenLocked events raised by the Main contract.
type MainTokenLockedIterator struct {
	Event *MainTokenLocked // Event containing the contract specifics and raw log

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
func (it *MainTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainTokenLocked)
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
		it.Event = new(MainTokenLocked)
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
func (it *MainTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainTokenLocked represents a TokenLocked event raised by the Main contract.
type MainTokenLocked struct {
	From   common.Address
	To     string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenLocked is a free log retrieval operation binding the contract event 0x865128631edc31d5501681b824f124e2f7718e3569d0bad93872617fce9f97f8.
//
// Solidity: event TokenLocked(address indexed from, string to, uint256 amount)
func (_Main *MainFilterer) FilterTokenLocked(opts *bind.FilterOpts, from []common.Address) (*MainTokenLockedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Main.contract.FilterLogs(opts, "TokenLocked", fromRule)
	if err != nil {
		return nil, err
	}
	return &MainTokenLockedIterator{contract: _Main.contract, event: "TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTokenLocked is a free log subscription operation binding the contract event 0x865128631edc31d5501681b824f124e2f7718e3569d0bad93872617fce9f97f8.
//
// Solidity: event TokenLocked(address indexed from, string to, uint256 amount)
func (_Main *MainFilterer) WatchTokenLocked(opts *bind.WatchOpts, sink chan<- *MainTokenLocked, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Main.contract.WatchLogs(opts, "TokenLocked", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainTokenLocked)
				if err := _Main.contract.UnpackLog(event, "TokenLocked", log); err != nil {
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
func (_Main *MainFilterer) ParseTokenLocked(log types.Log) (*MainTokenLocked, error) {
	event := new(MainTokenLocked)
	if err := _Main.contract.UnpackLog(event, "TokenLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}
