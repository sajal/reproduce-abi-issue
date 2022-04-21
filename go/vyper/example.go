// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vyper

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VyperMetaData contains all meta data concerning the Vyper contract.
var VyperMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"pure\",\"type\":\"function\",\"name\":\"doMath\",\"inputs\":[{\"name\":\"items\",\"type\":\"(uint256,uint256)[5]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2760}]",
}

// VyperABI is the input ABI used to generate the binding from.
// Deprecated: Use VyperMetaData.ABI instead.
var VyperABI = VyperMetaData.ABI

// Vyper is an auto generated Go binding around an Ethereum contract.
type Vyper struct {
	VyperCaller     // Read-only binding to the contract
	VyperTransactor // Write-only binding to the contract
	VyperFilterer   // Log filterer for contract events
}

// VyperCaller is an auto generated read-only Go binding around an Ethereum contract.
type VyperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VyperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VyperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VyperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VyperSession struct {
	Contract     *Vyper            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VyperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VyperCallerSession struct {
	Contract *VyperCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VyperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VyperTransactorSession struct {
	Contract     *VyperTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VyperRaw is an auto generated low-level Go binding around an Ethereum contract.
type VyperRaw struct {
	Contract *Vyper // Generic contract binding to access the raw methods on
}

// VyperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VyperCallerRaw struct {
	Contract *VyperCaller // Generic read-only contract binding to access the raw methods on
}

// VyperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VyperTransactorRaw struct {
	Contract *VyperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVyper creates a new instance of Vyper, bound to a specific deployed contract.
func NewVyper(address common.Address, backend bind.ContractBackend) (*Vyper, error) {
	contract, err := bindVyper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vyper{VyperCaller: VyperCaller{contract: contract}, VyperTransactor: VyperTransactor{contract: contract}, VyperFilterer: VyperFilterer{contract: contract}}, nil
}

// NewVyperCaller creates a new read-only instance of Vyper, bound to a specific deployed contract.
func NewVyperCaller(address common.Address, caller bind.ContractCaller) (*VyperCaller, error) {
	contract, err := bindVyper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VyperCaller{contract: contract}, nil
}

// NewVyperTransactor creates a new write-only instance of Vyper, bound to a specific deployed contract.
func NewVyperTransactor(address common.Address, transactor bind.ContractTransactor) (*VyperTransactor, error) {
	contract, err := bindVyper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VyperTransactor{contract: contract}, nil
}

// NewVyperFilterer creates a new log filterer instance of Vyper, bound to a specific deployed contract.
func NewVyperFilterer(address common.Address, filterer bind.ContractFilterer) (*VyperFilterer, error) {
	contract, err := bindVyper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VyperFilterer{contract: contract}, nil
}

// bindVyper binds a generic wrapper to an already deployed contract.
func bindVyper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VyperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vyper *VyperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vyper.Contract.VyperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vyper *VyperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vyper.Contract.VyperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vyper *VyperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vyper.Contract.VyperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vyper *VyperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vyper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vyper *VyperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vyper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vyper *VyperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vyper.Contract.contract.Transact(opts, method, params...)
}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func (_Vyper *VyperCaller) DoMath(opts *bind.CallOpts, items [5]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vyper.contract.Call(opts, &out, "doMath", items)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func (_Vyper *VyperSession) DoMath(items [5]*big.Int) (*big.Int, error) {
	return _Vyper.Contract.DoMath(&_Vyper.CallOpts, items)
}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func (_Vyper *VyperCallerSession) DoMath(items [5]*big.Int) (*big.Int, error) {
	return _Vyper.Contract.DoMath(&_Vyper.CallOpts, items)
}
