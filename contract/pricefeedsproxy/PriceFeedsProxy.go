// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pricefeedsproxy

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
	_ = abi.ConvertType
)

// PricefeedsproxyMetaData contains all meta data concerning the Pricefeedsproxy contract.
var PricefeedsproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b506040516107ab3803806107ab833981810160405281019061003391906100d7565b5b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5061015256610151565b6000815190506100d081610136565b5b92915050565b6000602082840312156100ea5760006000fd5b60006100f8848285016100c1565b9150505b92915050565b600061010d82610115565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b61013f81610102565b8114151561014d5760006000fd5b5b50565b5b61064a806101616000396000f3fe6080604052600436106100385760003560e01c80635c60da1b1461010757806383f94db714610133578063f2fde38b1461015d57610039565b5b5b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156100d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ca90610591565b60405180910390fd5b6040513660008237600060003683855af43d806000843e81600081146100fb578184f36100ff565b8184fd5b50505050505b005b3480156101145760006000fd5b5061011d610187565b60405161012a9190610533565b60405180910390f35b3480156101405760006000fd5b5061015b600480360381019061015691906103e9565b6101ad565b005b34801561016a5760006000fd5b50610185600480360381019061018091906103e9565b610285565b005b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561023f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102369061054f565b60405180910390fd5b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610317576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030e9061054f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610389576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038090610570565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b5056610613565b6000813590506103e2816105f8565b5b92915050565b6000602082840312156103fc5760006000fd5b600061040a848285016103d3565b9150505b92915050565b61041d816105c4565b82525b5050565b60006104316026836105b2565b91507f4f6e6c7920746865206f776e65722063616e2063616c6c20746869732066756e60008301527f6374696f6e2e000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006104986025836105b2565b91507f4e6577206f776e65722063616e6e6f7420626520746865207a65726f2061646460008301527f726573732e00000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006104ff6020836105b2565b91507f496d706c656d656e746174696f6e20636f6e7472616374206e6f74207365742e60008301526020820190505b919050565b60006020820190506105486000830184610414565b5b92915050565b6000602082019050818103600083015261056881610424565b90505b919050565b600060208201905081810360008301526105898161048b565b90505b919050565b600060208201905081810360008301526105aa816104f2565b90505b919050565b60008282526020820190505b92915050565b60006105cf826105d7565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b610601816105c4565b8114151561060f5760006000fd5b5b50565bfea264697066735822122032d91113a7fd5438c64caf0fcbb3c39c175208021cd05a423697aec0e59df8cf64736f6c63430008000033",
}

// PricefeedsproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use PricefeedsproxyMetaData.ABI instead.
var PricefeedsproxyABI = PricefeedsproxyMetaData.ABI

// PricefeedsproxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PricefeedsproxyMetaData.Bin instead.
var PricefeedsproxyBin = PricefeedsproxyMetaData.Bin

// DeployPricefeedsproxy deploys a new Ethereum contract, binding an instance of Pricefeedsproxy to it.
func DeployPricefeedsproxy(auth *bind.TransactOpts, backend bind.ContractBackend, _implementation common.Address) (common.Address, *types.Transaction, *Pricefeedsproxy, error) {
	parsed, err := PricefeedsproxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PricefeedsproxyBin), backend, _implementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pricefeedsproxy{PricefeedsproxyCaller: PricefeedsproxyCaller{contract: contract}, PricefeedsproxyTransactor: PricefeedsproxyTransactor{contract: contract}, PricefeedsproxyFilterer: PricefeedsproxyFilterer{contract: contract}}, nil
}

// Pricefeedsproxy is an auto generated Go binding around an Ethereum contract.
type Pricefeedsproxy struct {
	PricefeedsproxyCaller     // Read-only binding to the contract
	PricefeedsproxyTransactor // Write-only binding to the contract
	PricefeedsproxyFilterer   // Log filterer for contract events
}

// PricefeedsproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PricefeedsproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricefeedsproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PricefeedsproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricefeedsproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PricefeedsproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricefeedsproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PricefeedsproxySession struct {
	Contract     *Pricefeedsproxy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PricefeedsproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PricefeedsproxyCallerSession struct {
	Contract *PricefeedsproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PricefeedsproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PricefeedsproxyTransactorSession struct {
	Contract     *PricefeedsproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PricefeedsproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PricefeedsproxyRaw struct {
	Contract *Pricefeedsproxy // Generic contract binding to access the raw methods on
}

// PricefeedsproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PricefeedsproxyCallerRaw struct {
	Contract *PricefeedsproxyCaller // Generic read-only contract binding to access the raw methods on
}

// PricefeedsproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PricefeedsproxyTransactorRaw struct {
	Contract *PricefeedsproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPricefeedsproxy creates a new instance of Pricefeedsproxy, bound to a specific deployed contract.
func NewPricefeedsproxy(address common.Address, backend bind.ContractBackend) (*Pricefeedsproxy, error) {
	contract, err := bindPricefeedsproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pricefeedsproxy{PricefeedsproxyCaller: PricefeedsproxyCaller{contract: contract}, PricefeedsproxyTransactor: PricefeedsproxyTransactor{contract: contract}, PricefeedsproxyFilterer: PricefeedsproxyFilterer{contract: contract}}, nil
}

// NewPricefeedsproxyCaller creates a new read-only instance of Pricefeedsproxy, bound to a specific deployed contract.
func NewPricefeedsproxyCaller(address common.Address, caller bind.ContractCaller) (*PricefeedsproxyCaller, error) {
	contract, err := bindPricefeedsproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PricefeedsproxyCaller{contract: contract}, nil
}

// NewPricefeedsproxyTransactor creates a new write-only instance of Pricefeedsproxy, bound to a specific deployed contract.
func NewPricefeedsproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*PricefeedsproxyTransactor, error) {
	contract, err := bindPricefeedsproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PricefeedsproxyTransactor{contract: contract}, nil
}

// NewPricefeedsproxyFilterer creates a new log filterer instance of Pricefeedsproxy, bound to a specific deployed contract.
func NewPricefeedsproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*PricefeedsproxyFilterer, error) {
	contract, err := bindPricefeedsproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PricefeedsproxyFilterer{contract: contract}, nil
}

// bindPricefeedsproxy binds a generic wrapper to an already deployed contract.
func bindPricefeedsproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PricefeedsproxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricefeedsproxy *PricefeedsproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pricefeedsproxy.Contract.PricefeedsproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricefeedsproxy *PricefeedsproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.PricefeedsproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricefeedsproxy *PricefeedsproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.PricefeedsproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricefeedsproxy *PricefeedsproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pricefeedsproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricefeedsproxy *PricefeedsproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricefeedsproxy *PricefeedsproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Pricefeedsproxy *PricefeedsproxyCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pricefeedsproxy.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Pricefeedsproxy *PricefeedsproxySession) Implementation() (common.Address, error) {
	return _Pricefeedsproxy.Contract.Implementation(&_Pricefeedsproxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Pricefeedsproxy *PricefeedsproxyCallerSession) Implementation() (common.Address, error) {
	return _Pricefeedsproxy.Contract.Implementation(&_Pricefeedsproxy.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pricefeedsproxy *PricefeedsproxyTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Pricefeedsproxy.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pricefeedsproxy *PricefeedsproxySession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.TransferOwnership(&_Pricefeedsproxy.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Pricefeedsproxy *PricefeedsproxyTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.TransferOwnership(&_Pricefeedsproxy.TransactOpts, _newOwner)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newImplementation) returns()
func (_Pricefeedsproxy *PricefeedsproxyTransactor) UpgradeImplementation(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _Pricefeedsproxy.contract.Transact(opts, "upgradeImplementation", _newImplementation)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newImplementation) returns()
func (_Pricefeedsproxy *PricefeedsproxySession) UpgradeImplementation(_newImplementation common.Address) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.UpgradeImplementation(&_Pricefeedsproxy.TransactOpts, _newImplementation)
}

// UpgradeImplementation is a paid mutator transaction binding the contract method 0x83f94db7.
//
// Solidity: function upgradeImplementation(address _newImplementation) returns()
func (_Pricefeedsproxy *PricefeedsproxyTransactorSession) UpgradeImplementation(_newImplementation common.Address) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.UpgradeImplementation(&_Pricefeedsproxy.TransactOpts, _newImplementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Pricefeedsproxy *PricefeedsproxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Pricefeedsproxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Pricefeedsproxy *PricefeedsproxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.Fallback(&_Pricefeedsproxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Pricefeedsproxy *PricefeedsproxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Pricefeedsproxy.Contract.Fallback(&_Pricefeedsproxy.TransactOpts, calldata)
}
