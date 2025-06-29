// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ArenaTokenMetaData contains all meta data concerning the ArenaToken contract.
var ArenaTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"airdropAllocationPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecosystemAllocationPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecosystemAllocationPercentageUnlocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAllocationPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ArenaTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ArenaTokenMetaData.ABI instead.
var ArenaTokenABI = ArenaTokenMetaData.ABI

// ArenaToken is an auto generated Go binding around an Ethereum contract.
type ArenaToken struct {
	ArenaTokenCaller     // Read-only binding to the contract
	ArenaTokenTransactor // Write-only binding to the contract
	ArenaTokenFilterer   // Log filterer for contract events
}

// ArenaTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArenaTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArenaTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArenaTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArenaTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArenaTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArenaTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArenaTokenSession struct {
	Contract     *ArenaToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArenaTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArenaTokenCallerSession struct {
	Contract *ArenaTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArenaTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArenaTokenTransactorSession struct {
	Contract     *ArenaTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArenaTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArenaTokenRaw struct {
	Contract *ArenaToken // Generic contract binding to access the raw methods on
}

// ArenaTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArenaTokenCallerRaw struct {
	Contract *ArenaTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ArenaTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArenaTokenTransactorRaw struct {
	Contract *ArenaTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArenaToken creates a new instance of ArenaToken, bound to a specific deployed contract.
func NewArenaToken(address common.Address, backend bind.ContractBackend) (*ArenaToken, error) {
	contract, err := bindArenaToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArenaToken{ArenaTokenCaller: ArenaTokenCaller{contract: contract}, ArenaTokenTransactor: ArenaTokenTransactor{contract: contract}, ArenaTokenFilterer: ArenaTokenFilterer{contract: contract}}, nil
}

// NewArenaTokenCaller creates a new read-only instance of ArenaToken, bound to a specific deployed contract.
func NewArenaTokenCaller(address common.Address, caller bind.ContractCaller) (*ArenaTokenCaller, error) {
	contract, err := bindArenaToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArenaTokenCaller{contract: contract}, nil
}

// NewArenaTokenTransactor creates a new write-only instance of ArenaToken, bound to a specific deployed contract.
func NewArenaTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ArenaTokenTransactor, error) {
	contract, err := bindArenaToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArenaTokenTransactor{contract: contract}, nil
}

// NewArenaTokenFilterer creates a new log filterer instance of ArenaToken, bound to a specific deployed contract.
func NewArenaTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ArenaTokenFilterer, error) {
	contract, err := bindArenaToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArenaTokenFilterer{contract: contract}, nil
}

// bindArenaToken binds a generic wrapper to an already deployed contract.
func bindArenaToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArenaTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArenaToken *ArenaTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArenaToken.Contract.ArenaTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArenaToken *ArenaTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArenaToken.Contract.ArenaTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArenaToken *ArenaTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArenaToken.Contract.ArenaTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArenaToken *ArenaTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArenaToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArenaToken *ArenaTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArenaToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArenaToken *ArenaTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArenaToken.Contract.contract.Transact(opts, method, params...)
}

// AirdropAllocationPercentage is a free data retrieval call binding the contract method 0xa85085c9.
//
// Solidity: function airdropAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) AirdropAllocationPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "airdropAllocationPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirdropAllocationPercentage is a free data retrieval call binding the contract method 0xa85085c9.
//
// Solidity: function airdropAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenSession) AirdropAllocationPercentage() (*big.Int, error) {
	return _ArenaToken.Contract.AirdropAllocationPercentage(&_ArenaToken.CallOpts)
}

// AirdropAllocationPercentage is a free data retrieval call binding the contract method 0xa85085c9.
//
// Solidity: function airdropAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) AirdropAllocationPercentage() (*big.Int, error) {
	return _ArenaToken.Contract.AirdropAllocationPercentage(&_ArenaToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ArenaToken *ArenaTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ArenaToken.Contract.Allowance(&_ArenaToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ArenaToken.Contract.Allowance(&_ArenaToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ArenaToken *ArenaTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ArenaToken.Contract.BalanceOf(&_ArenaToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ArenaToken.Contract.BalanceOf(&_ArenaToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ArenaToken *ArenaTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ArenaToken *ArenaTokenSession) Decimals() (uint8, error) {
	return _ArenaToken.Contract.Decimals(&_ArenaToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_ArenaToken *ArenaTokenCallerSession) Decimals() (uint8, error) {
	return _ArenaToken.Contract.Decimals(&_ArenaToken.CallOpts)
}

// EcosystemAllocationPercentage is a free data retrieval call binding the contract method 0xc797d577.
//
// Solidity: function ecosystemAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) EcosystemAllocationPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "ecosystemAllocationPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EcosystemAllocationPercentage is a free data retrieval call binding the contract method 0xc797d577.
//
// Solidity: function ecosystemAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenSession) EcosystemAllocationPercentage() (*big.Int, error) {
	return _ArenaToken.Contract.EcosystemAllocationPercentage(&_ArenaToken.CallOpts)
}

// EcosystemAllocationPercentage is a free data retrieval call binding the contract method 0xc797d577.
//
// Solidity: function ecosystemAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) EcosystemAllocationPercentage() (*big.Int, error) {
	return _ArenaToken.Contract.EcosystemAllocationPercentage(&_ArenaToken.CallOpts)
}

// EcosystemAllocationPercentageUnlocked is a free data retrieval call binding the contract method 0x463b6519.
//
// Solidity: function ecosystemAllocationPercentageUnlocked() view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) EcosystemAllocationPercentageUnlocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "ecosystemAllocationPercentageUnlocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EcosystemAllocationPercentageUnlocked is a free data retrieval call binding the contract method 0x463b6519.
//
// Solidity: function ecosystemAllocationPercentageUnlocked() view returns(uint256)
func (_ArenaToken *ArenaTokenSession) EcosystemAllocationPercentageUnlocked() (*big.Int, error) {
	return _ArenaToken.Contract.EcosystemAllocationPercentageUnlocked(&_ArenaToken.CallOpts)
}

// EcosystemAllocationPercentageUnlocked is a free data retrieval call binding the contract method 0x463b6519.
//
// Solidity: function ecosystemAllocationPercentageUnlocked() view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) EcosystemAllocationPercentageUnlocked() (*big.Int, error) {
	return _ArenaToken.Contract.EcosystemAllocationPercentageUnlocked(&_ArenaToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArenaToken *ArenaTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArenaToken *ArenaTokenSession) Name() (string, error) {
	return _ArenaToken.Contract.Name(&_ArenaToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ArenaToken *ArenaTokenCallerSession) Name() (string, error) {
	return _ArenaToken.Contract.Name(&_ArenaToken.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_ArenaToken *ArenaTokenSession) Supply() (*big.Int, error) {
	return _ArenaToken.Contract.Supply(&_ArenaToken.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) Supply() (*big.Int, error) {
	return _ArenaToken.Contract.Supply(&_ArenaToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArenaToken *ArenaTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArenaToken *ArenaTokenSession) Symbol() (string, error) {
	return _ArenaToken.Contract.Symbol(&_ArenaToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ArenaToken *ArenaTokenCallerSession) Symbol() (string, error) {
	return _ArenaToken.Contract.Symbol(&_ArenaToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArenaToken *ArenaTokenSession) TotalSupply() (*big.Int, error) {
	return _ArenaToken.Contract.TotalSupply(&_ArenaToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ArenaToken.Contract.TotalSupply(&_ArenaToken.CallOpts)
}

// TreasuryAllocationPercentage is a free data retrieval call binding the contract method 0xade3c5d7.
//
// Solidity: function treasuryAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenCaller) TreasuryAllocationPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArenaToken.contract.Call(opts, &out, "treasuryAllocationPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryAllocationPercentage is a free data retrieval call binding the contract method 0xade3c5d7.
//
// Solidity: function treasuryAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenSession) TreasuryAllocationPercentage() (*big.Int, error) {
	return _ArenaToken.Contract.TreasuryAllocationPercentage(&_ArenaToken.CallOpts)
}

// TreasuryAllocationPercentage is a free data retrieval call binding the contract method 0xade3c5d7.
//
// Solidity: function treasuryAllocationPercentage() view returns(uint256)
func (_ArenaToken *ArenaTokenCallerSession) TreasuryAllocationPercentage() (*big.Int, error) {
	return _ArenaToken.Contract.TreasuryAllocationPercentage(&_ArenaToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.Approve(&_ArenaToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.Approve(&_ArenaToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ArenaToken *ArenaTokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ArenaToken *ArenaTokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.Burn(&_ArenaToken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ArenaToken *ArenaTokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.Burn(&_ArenaToken.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ArenaToken *ArenaTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.contract.Transact(opts, "burnFrom", account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ArenaToken *ArenaTokenSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.BurnFrom(&_ArenaToken.TransactOpts, account, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 value) returns()
func (_ArenaToken *ArenaTokenTransactorSession) BurnFrom(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.BurnFrom(&_ArenaToken.TransactOpts, account, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.Transfer(&_ArenaToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.Transfer(&_ArenaToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.TransferFrom(&_ArenaToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ArenaToken *ArenaTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ArenaToken.Contract.TransferFrom(&_ArenaToken.TransactOpts, from, to, value)
}

// ArenaTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ArenaToken contract.
type ArenaTokenApprovalIterator struct {
	Event *ArenaTokenApproval // Event containing the contract specifics and raw log

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
func (it *ArenaTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArenaTokenApproval)
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
		it.Event = new(ArenaTokenApproval)
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
func (it *ArenaTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArenaTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArenaTokenApproval represents a Approval event raised by the ArenaToken contract.
type ArenaTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ArenaToken *ArenaTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ArenaTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ArenaToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ArenaTokenApprovalIterator{contract: _ArenaToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ArenaToken *ArenaTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ArenaTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ArenaToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArenaTokenApproval)
				if err := _ArenaToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ArenaToken *ArenaTokenFilterer) ParseApproval(log types.Log) (*ArenaTokenApproval, error) {
	event := new(ArenaTokenApproval)
	if err := _ArenaToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArenaTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ArenaToken contract.
type ArenaTokenTransferIterator struct {
	Event *ArenaTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ArenaTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArenaTokenTransfer)
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
		it.Event = new(ArenaTokenTransfer)
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
func (it *ArenaTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArenaTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArenaTokenTransfer represents a Transfer event raised by the ArenaToken contract.
type ArenaTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ArenaToken *ArenaTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArenaTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArenaToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArenaTokenTransferIterator{contract: _ArenaToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ArenaToken *ArenaTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ArenaTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArenaToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArenaTokenTransfer)
				if err := _ArenaToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ArenaToken *ArenaTokenFilterer) ParseTransfer(log types.Log) (*ArenaTokenTransfer, error) {
	event := new(ArenaTokenTransfer)
	if err := _ArenaToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
