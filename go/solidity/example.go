// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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

// ExampleNumbers is an auto generated low-level Go binding around an user-defined struct.
type ExampleNumbers struct {
	A *big.Int
	B *big.Int
}

// SolidityMetaData contains all meta data concerning the Solidity contract.
var SolidityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"internalType\":\"structExample.Numbers[5]\",\"name\":\"items\",\"type\":\"tuple[5]\"}],\"name\":\"doMath\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SolidityABI is the input ABI used to generate the binding from.
// Deprecated: Use SolidityMetaData.ABI instead.
var SolidityABI = SolidityMetaData.ABI

// Solidity is an auto generated Go binding around an Ethereum contract.
type Solidity struct {
	SolidityCaller     // Read-only binding to the contract
	SolidityTransactor // Write-only binding to the contract
	SolidityFilterer   // Log filterer for contract events
}

// SolidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoliditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoliditySession struct {
	Contract     *Solidity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolidityCallerSession struct {
	Contract *SolidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SolidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolidityTransactorSession struct {
	Contract     *SolidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolidityRaw struct {
	Contract *Solidity // Generic contract binding to access the raw methods on
}

// SolidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolidityCallerRaw struct {
	Contract *SolidityCaller // Generic read-only contract binding to access the raw methods on
}

// SolidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolidityTransactorRaw struct {
	Contract *SolidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolidity creates a new instance of Solidity, bound to a specific deployed contract.
func NewSolidity(address common.Address, backend bind.ContractBackend) (*Solidity, error) {
	contract, err := bindSolidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// NewSolidityCaller creates a new read-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityCaller(address common.Address, caller bind.ContractCaller) (*SolidityCaller, error) {
	contract, err := bindSolidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityCaller{contract: contract}, nil
}

// NewSolidityTransactor creates a new write-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityTransactor(address common.Address, transactor bind.ContractTransactor) (*SolidityTransactor, error) {
	contract, err := bindSolidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityTransactor{contract: contract}, nil
}

// NewSolidityFilterer creates a new log filterer instance of Solidity, bound to a specific deployed contract.
func NewSolidityFilterer(address common.Address, filterer bind.ContractFilterer) (*SolidityFilterer, error) {
	contract, err := bindSolidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolidityFilterer{contract: contract}, nil
}

// bindSolidity binds a generic wrapper to an already deployed contract.
func bindSolidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolidityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.SolidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transact(opts, method, params...)
}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func (_Solidity *SolidityCaller) DoMath(opts *bind.CallOpts, items [5]ExampleNumbers) (*big.Int, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "doMath", items)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func (_Solidity *SoliditySession) DoMath(items [5]ExampleNumbers) (*big.Int, error) {
	return _Solidity.Contract.DoMath(&_Solidity.CallOpts, items)
}

// DoMath is a free data retrieval call binding the contract method 0x0b48636c.
//
// Solidity: function doMath((uint256,uint256)[5] items) pure returns(uint256)
func (_Solidity *SolidityCallerSession) DoMath(items [5]ExampleNumbers) (*big.Int, error) {
	return _Solidity.Contract.DoMath(&_Solidity.CallOpts, items)
}
