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

// RouterContractMetaData contains all meta data concerning the RouterContract contract.
var RouterContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WAVAX\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WAVAX\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityAVAXSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAVAXMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAVAX\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAXSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterContractMetaData.ABI instead.
var RouterContractABI = RouterContractMetaData.ABI

// RouterContract is an auto generated Go binding around an Ethereum contract.
type RouterContract struct {
	RouterContractCaller     // Read-only binding to the contract
	RouterContractTransactor // Write-only binding to the contract
	RouterContractFilterer   // Log filterer for contract events
}

// RouterContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterContractSession struct {
	Contract     *RouterContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterContractCallerSession struct {
	Contract *RouterContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RouterContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterContractTransactorSession struct {
	Contract     *RouterContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RouterContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterContractRaw struct {
	Contract *RouterContract // Generic contract binding to access the raw methods on
}

// RouterContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterContractCallerRaw struct {
	Contract *RouterContractCaller // Generic read-only contract binding to access the raw methods on
}

// RouterContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterContractTransactorRaw struct {
	Contract *RouterContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterContract creates a new instance of RouterContract, bound to a specific deployed contract.
func NewRouterContract(address common.Address, backend bind.ContractBackend) (*RouterContract, error) {
	contract, err := bindRouterContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RouterContract{RouterContractCaller: RouterContractCaller{contract: contract}, RouterContractTransactor: RouterContractTransactor{contract: contract}, RouterContractFilterer: RouterContractFilterer{contract: contract}}, nil
}

// NewRouterContractCaller creates a new read-only instance of RouterContract, bound to a specific deployed contract.
func NewRouterContractCaller(address common.Address, caller bind.ContractCaller) (*RouterContractCaller, error) {
	contract, err := bindRouterContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterContractCaller{contract: contract}, nil
}

// NewRouterContractTransactor creates a new write-only instance of RouterContract, bound to a specific deployed contract.
func NewRouterContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterContractTransactor, error) {
	contract, err := bindRouterContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterContractTransactor{contract: contract}, nil
}

// NewRouterContractFilterer creates a new log filterer instance of RouterContract, bound to a specific deployed contract.
func NewRouterContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterContractFilterer, error) {
	contract, err := bindRouterContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterContractFilterer{contract: contract}, nil
}

// bindRouterContract binds a generic wrapper to an already deployed contract.
func bindRouterContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterContract *RouterContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterContract.Contract.RouterContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterContract *RouterContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterContract.Contract.RouterContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterContract *RouterContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterContract.Contract.RouterContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RouterContract *RouterContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RouterContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RouterContract *RouterContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RouterContract *RouterContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RouterContract.Contract.contract.Transact(opts, method, params...)
}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (_RouterContract *RouterContractCaller) WAVAX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "WAVAX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (_RouterContract *RouterContractSession) WAVAX() (common.Address, error) {
	return _RouterContract.Contract.WAVAX(&_RouterContract.CallOpts)
}

// WAVAX is a free data retrieval call binding the contract method 0x73b295c2.
//
// Solidity: function WAVAX() view returns(address)
func (_RouterContract *RouterContractCallerSession) WAVAX() (common.Address, error) {
	return _RouterContract.Contract.WAVAX(&_RouterContract.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_RouterContract *RouterContractCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_RouterContract *RouterContractSession) Factory() (common.Address, error) {
	return _RouterContract.Contract.Factory(&_RouterContract.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_RouterContract *RouterContractCallerSession) Factory() (common.Address, error) {
	return _RouterContract.Contract.Factory(&_RouterContract.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_RouterContract *RouterContractCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_RouterContract *RouterContractSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _RouterContract.Contract.GetAmountIn(&_RouterContract.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_RouterContract *RouterContractCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _RouterContract.Contract.GetAmountIn(&_RouterContract.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_RouterContract *RouterContractCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_RouterContract *RouterContractSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _RouterContract.Contract.GetAmountOut(&_RouterContract.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_RouterContract *RouterContractCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _RouterContract.Contract.GetAmountOut(&_RouterContract.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_RouterContract *RouterContractCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _RouterContract.Contract.GetAmountsIn(&_RouterContract.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_RouterContract *RouterContractCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _RouterContract.Contract.GetAmountsIn(&_RouterContract.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_RouterContract *RouterContractCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _RouterContract.Contract.GetAmountsOut(&_RouterContract.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_RouterContract *RouterContractCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _RouterContract.Contract.GetAmountsOut(&_RouterContract.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_RouterContract *RouterContractCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RouterContract.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_RouterContract *RouterContractSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _RouterContract.Contract.Quote(&_RouterContract.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_RouterContract *RouterContractCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _RouterContract.Contract.Quote(&_RouterContract.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_RouterContract *RouterContractTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_RouterContract *RouterContractSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.AddLiquidity(&_RouterContract.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_RouterContract *RouterContractTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.AddLiquidity(&_RouterContract.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_RouterContract *RouterContractTransactor) AddLiquidityAVAX(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "addLiquidityAVAX", token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_RouterContract *RouterContractSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.AddLiquidityAVAX(&_RouterContract.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// AddLiquidityAVAX is a paid mutator transaction binding the contract method 0xf91b3f72.
//
// Solidity: function addLiquidityAVAX(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountAVAX, uint256 liquidity)
func (_RouterContract *RouterContractTransactorSession) AddLiquidityAVAX(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.AddLiquidityAVAX(&_RouterContract.TransactOpts, token, amountTokenDesired, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_RouterContract *RouterContractTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_RouterContract *RouterContractSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidity(&_RouterContract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_RouterContract *RouterContractTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidity(&_RouterContract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_RouterContract *RouterContractTransactor) RemoveLiquidityAVAX(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "removeLiquidityAVAX", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_RouterContract *RouterContractSession) RemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAX(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAX is a paid mutator transaction binding the contract method 0x33c6b725.
//
// Solidity: function removeLiquidityAVAX(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountAVAX)
func (_RouterContract *RouterContractTransactorSession) RemoveLiquidityAVAX(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAX(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_RouterContract *RouterContractTransactor) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "removeLiquidityAVAXSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_RouterContract *RouterContractSession) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAXSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x73bc79cf.
//
// Solidity: function removeLiquidityAVAXSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline) returns(uint256 amountAVAX)
func (_RouterContract *RouterContractTransactorSession) RemoveLiquidityAVAXSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAXSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_RouterContract *RouterContractTransactor) RemoveLiquidityAVAXWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "removeLiquidityAVAXWithPermit", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_RouterContract *RouterContractSession) RemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAXWithPermit(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermit is a paid mutator transaction binding the contract method 0x2c407024.
//
// Solidity: function removeLiquidityAVAXWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountAVAX)
func (_RouterContract *RouterContractTransactorSession) RemoveLiquidityAVAXWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAXWithPermit(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_RouterContract *RouterContractTransactor) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_RouterContract *RouterContractSession) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x9fc27226.
//
// Solidity: function removeLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountAVAXMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountAVAX)
func (_RouterContract *RouterContractTransactorSession) RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountAVAXMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityAVAXWithPermitSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, token, liquidity, amountTokenMin, amountAVAXMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_RouterContract *RouterContractTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_RouterContract *RouterContractSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityWithPermit(&_RouterContract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_RouterContract *RouterContractTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _RouterContract.Contract.RemoveLiquidityWithPermit(&_RouterContract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapAVAXForExactTokens", amountOut, path, to, deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) SwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapAVAXForExactTokens(&_RouterContract.TransactOpts, amountOut, path, to, deadline)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactorSession) SwapAVAXForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapAVAXForExactTokens(&_RouterContract.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapExactAVAXForTokens", amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) SwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactAVAXForTokens(&_RouterContract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactorSession) SwapExactAVAXForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactAVAXForTokens(&_RouterContract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_RouterContract *RouterContractTransactor) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapExactAVAXForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_RouterContract *RouterContractSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactAVAXForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xc57559dd.
//
// Solidity: function swapExactAVAXForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_RouterContract *RouterContractTransactorSession) SwapExactAVAXForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactAVAXForTokensSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapExactTokensForAVAX", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) SwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForAVAX(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactorSession) SwapExactTokensForAVAX(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForAVAX(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_RouterContract *RouterContractTransactor) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapExactTokensForAVAXSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_RouterContract *RouterContractSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForAVAXSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x762b1562.
//
// Solidity: function swapExactTokensForAVAXSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_RouterContract *RouterContractTransactorSession) SwapExactTokensForAVAXSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForAVAXSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForTokens(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForTokens(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_RouterContract *RouterContractTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_RouterContract *RouterContractSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_RouterContract *RouterContractTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_RouterContract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapTokensForExactAVAX", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) SwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapTokensForExactAVAX(&_RouterContract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactorSession) SwapTokensForExactAVAX(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapTokensForExactAVAX(&_RouterContract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapTokensForExactTokens(&_RouterContract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_RouterContract *RouterContractTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _RouterContract.Contract.SwapTokensForExactTokens(&_RouterContract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterContract *RouterContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RouterContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterContract *RouterContractSession) Receive() (*types.Transaction, error) {
	return _RouterContract.Contract.Receive(&_RouterContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RouterContract *RouterContractTransactorSession) Receive() (*types.Transaction, error) {
	return _RouterContract.Contract.Receive(&_RouterContract.TransactOpts)
}
