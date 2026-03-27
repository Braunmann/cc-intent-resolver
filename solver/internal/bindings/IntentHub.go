// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// Intent is an auto generated low-level Go binding around an user-defined struct.
type Intent struct {
	Id              [32]byte
	Maker           common.Address
	InputToken      common.Address
	InputAmount     *big.Int
	OutputToken     common.Address
	Solver          common.Address
	MinOutputAmount *big.Int
	TargetChainId   uint32
	Recipient       common.Address
	Deadline        uint64
	Nonce           uint64
	Status          uint8
}

// IntentHubMetaData contains all meta data concerning the IntentHub contract.
var IntentHubMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelIntent\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createIntent\",\"inputs\":[{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minOutputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fulfillIntent\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getIntent\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIntent\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minOutputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIntentStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"intents\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minOutputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deadline\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIntentStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"settleIntent\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"IntentCancelled\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntentCreated\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minOutputAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"targetChainId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"deadline\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntentFulfilled\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntentSettled\",\"inputs\":[{\"name\":\"intentId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DeadlineInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IntentNotCancellable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IntentNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IntentNotFulfillable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IntentNotSettleable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotIntentMaker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotIntentSolver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
}

// IntentHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IntentHubMetaData.ABI instead.
var IntentHubABI = IntentHubMetaData.ABI

// IntentHub is an auto generated Go binding around an Ethereum contract.
type IntentHub struct {
	IntentHubCaller     // Read-only binding to the contract
	IntentHubTransactor // Write-only binding to the contract
	IntentHubFilterer   // Log filterer for contract events
}

// IntentHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IntentHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IntentHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IntentHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IntentHubSession struct {
	Contract     *IntentHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IntentHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IntentHubCallerSession struct {
	Contract *IntentHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IntentHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IntentHubTransactorSession struct {
	Contract     *IntentHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IntentHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IntentHubRaw struct {
	Contract *IntentHub // Generic contract binding to access the raw methods on
}

// IntentHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IntentHubCallerRaw struct {
	Contract *IntentHubCaller // Generic read-only contract binding to access the raw methods on
}

// IntentHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IntentHubTransactorRaw struct {
	Contract *IntentHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIntentHub creates a new instance of IntentHub, bound to a specific deployed contract.
func NewIntentHub(address common.Address, backend bind.ContractBackend) (*IntentHub, error) {
	contract, err := bindIntentHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IntentHub{IntentHubCaller: IntentHubCaller{contract: contract}, IntentHubTransactor: IntentHubTransactor{contract: contract}, IntentHubFilterer: IntentHubFilterer{contract: contract}}, nil
}

// NewIntentHubCaller creates a new read-only instance of IntentHub, bound to a specific deployed contract.
func NewIntentHubCaller(address common.Address, caller bind.ContractCaller) (*IntentHubCaller, error) {
	contract, err := bindIntentHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IntentHubCaller{contract: contract}, nil
}

// NewIntentHubTransactor creates a new write-only instance of IntentHub, bound to a specific deployed contract.
func NewIntentHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IntentHubTransactor, error) {
	contract, err := bindIntentHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IntentHubTransactor{contract: contract}, nil
}

// NewIntentHubFilterer creates a new log filterer instance of IntentHub, bound to a specific deployed contract.
func NewIntentHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IntentHubFilterer, error) {
	contract, err := bindIntentHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IntentHubFilterer{contract: contract}, nil
}

// bindIntentHub binds a generic wrapper to an already deployed contract.
func bindIntentHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IntentHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntentHub *IntentHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntentHub.Contract.IntentHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntentHub *IntentHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentHub.Contract.IntentHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntentHub *IntentHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntentHub.Contract.IntentHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntentHub *IntentHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntentHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntentHub *IntentHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntentHub *IntentHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntentHub.Contract.contract.Transact(opts, method, params...)
}

// GetIntent is a free data retrieval call binding the contract method 0xf13c46aa.
//
// Solidity: function getIntent(bytes32 intentId) view returns((bytes32,address,address,uint256,address,address,uint256,uint32,address,uint64,uint64,uint8))
func (_IntentHub *IntentHubCaller) GetIntent(opts *bind.CallOpts, intentId [32]byte) (Intent, error) {
	var out []interface{}
	err := _IntentHub.contract.Call(opts, &out, "getIntent", intentId)

	if err != nil {
		return *new(Intent), err
	}

	out0 := *abi.ConvertType(out[0], new(Intent)).(*Intent)

	return out0, err

}

// GetIntent is a free data retrieval call binding the contract method 0xf13c46aa.
//
// Solidity: function getIntent(bytes32 intentId) view returns((bytes32,address,address,uint256,address,address,uint256,uint32,address,uint64,uint64,uint8))
func (_IntentHub *IntentHubSession) GetIntent(intentId [32]byte) (Intent, error) {
	return _IntentHub.Contract.GetIntent(&_IntentHub.CallOpts, intentId)
}

// GetIntent is a free data retrieval call binding the contract method 0xf13c46aa.
//
// Solidity: function getIntent(bytes32 intentId) view returns((bytes32,address,address,uint256,address,address,uint256,uint32,address,uint64,uint64,uint8))
func (_IntentHub *IntentHubCallerSession) GetIntent(intentId [32]byte) (Intent, error) {
	return _IntentHub.Contract.GetIntent(&_IntentHub.CallOpts, intentId)
}

// Intents is a free data retrieval call binding the contract method 0x9021578a.
//
// Solidity: function intents(bytes32 ) view returns(bytes32 id, address maker, address inputToken, uint256 inputAmount, address outputToken, address solver, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline, uint64 nonce, uint8 status)
func (_IntentHub *IntentHubCaller) Intents(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Id              [32]byte
	Maker           common.Address
	InputToken      common.Address
	InputAmount     *big.Int
	OutputToken     common.Address
	Solver          common.Address
	MinOutputAmount *big.Int
	TargetChainId   uint32
	Recipient       common.Address
	Deadline        uint64
	Nonce           uint64
	Status          uint8
}, error) {
	var out []interface{}
	err := _IntentHub.contract.Call(opts, &out, "intents", arg0)

	outstruct := new(struct {
		Id              [32]byte
		Maker           common.Address
		InputToken      common.Address
		InputAmount     *big.Int
		OutputToken     common.Address
		Solver          common.Address
		MinOutputAmount *big.Int
		TargetChainId   uint32
		Recipient       common.Address
		Deadline        uint64
		Nonce           uint64
		Status          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Maker = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.InputToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.InputAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OutputToken = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Solver = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.MinOutputAmount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TargetChainId = *abi.ConvertType(out[7], new(uint32)).(*uint32)
	outstruct.Recipient = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.Deadline = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.Nonce = *abi.ConvertType(out[10], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[11], new(uint8)).(*uint8)

	return *outstruct, err

}

// Intents is a free data retrieval call binding the contract method 0x9021578a.
//
// Solidity: function intents(bytes32 ) view returns(bytes32 id, address maker, address inputToken, uint256 inputAmount, address outputToken, address solver, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline, uint64 nonce, uint8 status)
func (_IntentHub *IntentHubSession) Intents(arg0 [32]byte) (struct {
	Id              [32]byte
	Maker           common.Address
	InputToken      common.Address
	InputAmount     *big.Int
	OutputToken     common.Address
	Solver          common.Address
	MinOutputAmount *big.Int
	TargetChainId   uint32
	Recipient       common.Address
	Deadline        uint64
	Nonce           uint64
	Status          uint8
}, error) {
	return _IntentHub.Contract.Intents(&_IntentHub.CallOpts, arg0)
}

// Intents is a free data retrieval call binding the contract method 0x9021578a.
//
// Solidity: function intents(bytes32 ) view returns(bytes32 id, address maker, address inputToken, uint256 inputAmount, address outputToken, address solver, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline, uint64 nonce, uint8 status)
func (_IntentHub *IntentHubCallerSession) Intents(arg0 [32]byte) (struct {
	Id              [32]byte
	Maker           common.Address
	InputToken      common.Address
	InputAmount     *big.Int
	OutputToken     common.Address
	Solver          common.Address
	MinOutputAmount *big.Int
	TargetChainId   uint32
	Recipient       common.Address
	Deadline        uint64
	Nonce           uint64
	Status          uint8
}, error) {
	return _IntentHub.Contract.Intents(&_IntentHub.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint64)
func (_IntentHub *IntentHubCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _IntentHub.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint64)
func (_IntentHub *IntentHubSession) Nonces(arg0 common.Address) (uint64, error) {
	return _IntentHub.Contract.Nonces(&_IntentHub.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint64)
func (_IntentHub *IntentHubCallerSession) Nonces(arg0 common.Address) (uint64, error) {
	return _IntentHub.Contract.Nonces(&_IntentHub.CallOpts, arg0)
}

// CancelIntent is a paid mutator transaction binding the contract method 0xd55f960d.
//
// Solidity: function cancelIntent(bytes32 intentId) returns()
func (_IntentHub *IntentHubTransactor) CancelIntent(opts *bind.TransactOpts, intentId [32]byte) (*types.Transaction, error) {
	return _IntentHub.contract.Transact(opts, "cancelIntent", intentId)
}

// CancelIntent is a paid mutator transaction binding the contract method 0xd55f960d.
//
// Solidity: function cancelIntent(bytes32 intentId) returns()
func (_IntentHub *IntentHubSession) CancelIntent(intentId [32]byte) (*types.Transaction, error) {
	return _IntentHub.Contract.CancelIntent(&_IntentHub.TransactOpts, intentId)
}

// CancelIntent is a paid mutator transaction binding the contract method 0xd55f960d.
//
// Solidity: function cancelIntent(bytes32 intentId) returns()
func (_IntentHub *IntentHubTransactorSession) CancelIntent(intentId [32]byte) (*types.Transaction, error) {
	return _IntentHub.Contract.CancelIntent(&_IntentHub.TransactOpts, intentId)
}

// CreateIntent is a paid mutator transaction binding the contract method 0x24e14caf.
//
// Solidity: function createIntent(address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline) returns(bytes32)
func (_IntentHub *IntentHubTransactor) CreateIntent(opts *bind.TransactOpts, inputToken common.Address, inputAmount *big.Int, outputToken common.Address, minOutputAmount *big.Int, targetChainId uint32, recipient common.Address, deadline uint64) (*types.Transaction, error) {
	return _IntentHub.contract.Transact(opts, "createIntent", inputToken, inputAmount, outputToken, minOutputAmount, targetChainId, recipient, deadline)
}

// CreateIntent is a paid mutator transaction binding the contract method 0x24e14caf.
//
// Solidity: function createIntent(address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline) returns(bytes32)
func (_IntentHub *IntentHubSession) CreateIntent(inputToken common.Address, inputAmount *big.Int, outputToken common.Address, minOutputAmount *big.Int, targetChainId uint32, recipient common.Address, deadline uint64) (*types.Transaction, error) {
	return _IntentHub.Contract.CreateIntent(&_IntentHub.TransactOpts, inputToken, inputAmount, outputToken, minOutputAmount, targetChainId, recipient, deadline)
}

// CreateIntent is a paid mutator transaction binding the contract method 0x24e14caf.
//
// Solidity: function createIntent(address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline) returns(bytes32)
func (_IntentHub *IntentHubTransactorSession) CreateIntent(inputToken common.Address, inputAmount *big.Int, outputToken common.Address, minOutputAmount *big.Int, targetChainId uint32, recipient common.Address, deadline uint64) (*types.Transaction, error) {
	return _IntentHub.Contract.CreateIntent(&_IntentHub.TransactOpts, inputToken, inputAmount, outputToken, minOutputAmount, targetChainId, recipient, deadline)
}

// FulfillIntent is a paid mutator transaction binding the contract method 0x2b15ece0.
//
// Solidity: function fulfillIntent(bytes32 intentId, address solver) returns()
func (_IntentHub *IntentHubTransactor) FulfillIntent(opts *bind.TransactOpts, intentId [32]byte, solver common.Address) (*types.Transaction, error) {
	return _IntentHub.contract.Transact(opts, "fulfillIntent", intentId, solver)
}

// FulfillIntent is a paid mutator transaction binding the contract method 0x2b15ece0.
//
// Solidity: function fulfillIntent(bytes32 intentId, address solver) returns()
func (_IntentHub *IntentHubSession) FulfillIntent(intentId [32]byte, solver common.Address) (*types.Transaction, error) {
	return _IntentHub.Contract.FulfillIntent(&_IntentHub.TransactOpts, intentId, solver)
}

// FulfillIntent is a paid mutator transaction binding the contract method 0x2b15ece0.
//
// Solidity: function fulfillIntent(bytes32 intentId, address solver) returns()
func (_IntentHub *IntentHubTransactorSession) FulfillIntent(intentId [32]byte, solver common.Address) (*types.Transaction, error) {
	return _IntentHub.Contract.FulfillIntent(&_IntentHub.TransactOpts, intentId, solver)
}

// SettleIntent is a paid mutator transaction binding the contract method 0x7bf8bb88.
//
// Solidity: function settleIntent(bytes32 intentId) returns()
func (_IntentHub *IntentHubTransactor) SettleIntent(opts *bind.TransactOpts, intentId [32]byte) (*types.Transaction, error) {
	return _IntentHub.contract.Transact(opts, "settleIntent", intentId)
}

// SettleIntent is a paid mutator transaction binding the contract method 0x7bf8bb88.
//
// Solidity: function settleIntent(bytes32 intentId) returns()
func (_IntentHub *IntentHubSession) SettleIntent(intentId [32]byte) (*types.Transaction, error) {
	return _IntentHub.Contract.SettleIntent(&_IntentHub.TransactOpts, intentId)
}

// SettleIntent is a paid mutator transaction binding the contract method 0x7bf8bb88.
//
// Solidity: function settleIntent(bytes32 intentId) returns()
func (_IntentHub *IntentHubTransactorSession) SettleIntent(intentId [32]byte) (*types.Transaction, error) {
	return _IntentHub.Contract.SettleIntent(&_IntentHub.TransactOpts, intentId)
}

// IntentHubIntentCancelledIterator is returned from FilterIntentCancelled and is used to iterate over the raw logs and unpacked data for IntentCancelled events raised by the IntentHub contract.
type IntentHubIntentCancelledIterator struct {
	Event *IntentHubIntentCancelled // Event containing the contract specifics and raw log

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
func (it *IntentHubIntentCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentHubIntentCancelled)
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
		it.Event = new(IntentHubIntentCancelled)
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
func (it *IntentHubIntentCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentHubIntentCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentHubIntentCancelled represents a IntentCancelled event raised by the IntentHub contract.
type IntentHubIntentCancelled struct {
	IntentId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIntentCancelled is a free log retrieval operation binding the contract event 0xc08eb64db16a39d2848960af04e3f16fb404d9d436a9f0e9d7d0d4854715c9dc.
//
// Solidity: event IntentCancelled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) FilterIntentCancelled(opts *bind.FilterOpts, intentId [][32]byte) (*IntentHubIntentCancelledIterator, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}

	logs, sub, err := _IntentHub.contract.FilterLogs(opts, "IntentCancelled", intentIdRule)
	if err != nil {
		return nil, err
	}
	return &IntentHubIntentCancelledIterator{contract: _IntentHub.contract, event: "IntentCancelled", logs: logs, sub: sub}, nil
}

// WatchIntentCancelled is a free log subscription operation binding the contract event 0xc08eb64db16a39d2848960af04e3f16fb404d9d436a9f0e9d7d0d4854715c9dc.
//
// Solidity: event IntentCancelled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) WatchIntentCancelled(opts *bind.WatchOpts, sink chan<- *IntentHubIntentCancelled, intentId [][32]byte) (event.Subscription, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}

	logs, sub, err := _IntentHub.contract.WatchLogs(opts, "IntentCancelled", intentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentHubIntentCancelled)
				if err := _IntentHub.contract.UnpackLog(event, "IntentCancelled", log); err != nil {
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

// ParseIntentCancelled is a log parse operation binding the contract event 0xc08eb64db16a39d2848960af04e3f16fb404d9d436a9f0e9d7d0d4854715c9dc.
//
// Solidity: event IntentCancelled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) ParseIntentCancelled(log types.Log) (*IntentHubIntentCancelled, error) {
	event := new(IntentHubIntentCancelled)
	if err := _IntentHub.contract.UnpackLog(event, "IntentCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentHubIntentCreatedIterator is returned from FilterIntentCreated and is used to iterate over the raw logs and unpacked data for IntentCreated events raised by the IntentHub contract.
type IntentHubIntentCreatedIterator struct {
	Event *IntentHubIntentCreated // Event containing the contract specifics and raw log

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
func (it *IntentHubIntentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentHubIntentCreated)
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
		it.Event = new(IntentHubIntentCreated)
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
func (it *IntentHubIntentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentHubIntentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentHubIntentCreated represents a IntentCreated event raised by the IntentHub contract.
type IntentHubIntentCreated struct {
	IntentId        [32]byte
	Maker           common.Address
	InputToken      common.Address
	InputAmount     *big.Int
	OutputToken     common.Address
	MinOutputAmount *big.Int
	TargetChainId   uint32
	Recipient       common.Address
	Deadline        uint64
	Nonce           uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIntentCreated is a free log retrieval operation binding the contract event 0x449f01145bea5d084f07cc1040bef141f0cb5ae7dbd57b824f670914d6af68e2.
//
// Solidity: event IntentCreated(bytes32 indexed intentId, address indexed maker, address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline, uint64 nonce)
func (_IntentHub *IntentHubFilterer) FilterIntentCreated(opts *bind.FilterOpts, intentId [][32]byte, maker []common.Address) (*IntentHubIntentCreatedIterator, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _IntentHub.contract.FilterLogs(opts, "IntentCreated", intentIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return &IntentHubIntentCreatedIterator{contract: _IntentHub.contract, event: "IntentCreated", logs: logs, sub: sub}, nil
}

// WatchIntentCreated is a free log subscription operation binding the contract event 0x449f01145bea5d084f07cc1040bef141f0cb5ae7dbd57b824f670914d6af68e2.
//
// Solidity: event IntentCreated(bytes32 indexed intentId, address indexed maker, address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline, uint64 nonce)
func (_IntentHub *IntentHubFilterer) WatchIntentCreated(opts *bind.WatchOpts, sink chan<- *IntentHubIntentCreated, intentId [][32]byte, maker []common.Address) (event.Subscription, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _IntentHub.contract.WatchLogs(opts, "IntentCreated", intentIdRule, makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentHubIntentCreated)
				if err := _IntentHub.contract.UnpackLog(event, "IntentCreated", log); err != nil {
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

// ParseIntentCreated is a log parse operation binding the contract event 0x449f01145bea5d084f07cc1040bef141f0cb5ae7dbd57b824f670914d6af68e2.
//
// Solidity: event IntentCreated(bytes32 indexed intentId, address indexed maker, address inputToken, uint256 inputAmount, address outputToken, uint256 minOutputAmount, uint32 targetChainId, address recipient, uint64 deadline, uint64 nonce)
func (_IntentHub *IntentHubFilterer) ParseIntentCreated(log types.Log) (*IntentHubIntentCreated, error) {
	event := new(IntentHubIntentCreated)
	if err := _IntentHub.contract.UnpackLog(event, "IntentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentHubIntentFulfilledIterator is returned from FilterIntentFulfilled and is used to iterate over the raw logs and unpacked data for IntentFulfilled events raised by the IntentHub contract.
type IntentHubIntentFulfilledIterator struct {
	Event *IntentHubIntentFulfilled // Event containing the contract specifics and raw log

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
func (it *IntentHubIntentFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentHubIntentFulfilled)
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
		it.Event = new(IntentHubIntentFulfilled)
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
func (it *IntentHubIntentFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentHubIntentFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentHubIntentFulfilled represents a IntentFulfilled event raised by the IntentHub contract.
type IntentHubIntentFulfilled struct {
	IntentId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIntentFulfilled is a free log retrieval operation binding the contract event 0x0c18079d729f8e6d119764c23f720af06195b6fd66cb2fcb021c049e2c6a86f4.
//
// Solidity: event IntentFulfilled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) FilterIntentFulfilled(opts *bind.FilterOpts, intentId [][32]byte) (*IntentHubIntentFulfilledIterator, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}

	logs, sub, err := _IntentHub.contract.FilterLogs(opts, "IntentFulfilled", intentIdRule)
	if err != nil {
		return nil, err
	}
	return &IntentHubIntentFulfilledIterator{contract: _IntentHub.contract, event: "IntentFulfilled", logs: logs, sub: sub}, nil
}

// WatchIntentFulfilled is a free log subscription operation binding the contract event 0x0c18079d729f8e6d119764c23f720af06195b6fd66cb2fcb021c049e2c6a86f4.
//
// Solidity: event IntentFulfilled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) WatchIntentFulfilled(opts *bind.WatchOpts, sink chan<- *IntentHubIntentFulfilled, intentId [][32]byte) (event.Subscription, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}

	logs, sub, err := _IntentHub.contract.WatchLogs(opts, "IntentFulfilled", intentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentHubIntentFulfilled)
				if err := _IntentHub.contract.UnpackLog(event, "IntentFulfilled", log); err != nil {
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

// ParseIntentFulfilled is a log parse operation binding the contract event 0x0c18079d729f8e6d119764c23f720af06195b6fd66cb2fcb021c049e2c6a86f4.
//
// Solidity: event IntentFulfilled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) ParseIntentFulfilled(log types.Log) (*IntentHubIntentFulfilled, error) {
	event := new(IntentHubIntentFulfilled)
	if err := _IntentHub.contract.UnpackLog(event, "IntentFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentHubIntentSettledIterator is returned from FilterIntentSettled and is used to iterate over the raw logs and unpacked data for IntentSettled events raised by the IntentHub contract.
type IntentHubIntentSettledIterator struct {
	Event *IntentHubIntentSettled // Event containing the contract specifics and raw log

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
func (it *IntentHubIntentSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentHubIntentSettled)
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
		it.Event = new(IntentHubIntentSettled)
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
func (it *IntentHubIntentSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentHubIntentSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentHubIntentSettled represents a IntentSettled event raised by the IntentHub contract.
type IntentHubIntentSettled struct {
	IntentId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIntentSettled is a free log retrieval operation binding the contract event 0x3d59f7aa41c3b6ea14d68eb6d4fb841f9160a9b21ef5dca3b617d1d34eb712d4.
//
// Solidity: event IntentSettled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) FilterIntentSettled(opts *bind.FilterOpts, intentId [][32]byte) (*IntentHubIntentSettledIterator, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}

	logs, sub, err := _IntentHub.contract.FilterLogs(opts, "IntentSettled", intentIdRule)
	if err != nil {
		return nil, err
	}
	return &IntentHubIntentSettledIterator{contract: _IntentHub.contract, event: "IntentSettled", logs: logs, sub: sub}, nil
}

// WatchIntentSettled is a free log subscription operation binding the contract event 0x3d59f7aa41c3b6ea14d68eb6d4fb841f9160a9b21ef5dca3b617d1d34eb712d4.
//
// Solidity: event IntentSettled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) WatchIntentSettled(opts *bind.WatchOpts, sink chan<- *IntentHubIntentSettled, intentId [][32]byte) (event.Subscription, error) {

	var intentIdRule []interface{}
	for _, intentIdItem := range intentId {
		intentIdRule = append(intentIdRule, intentIdItem)
	}

	logs, sub, err := _IntentHub.contract.WatchLogs(opts, "IntentSettled", intentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentHubIntentSettled)
				if err := _IntentHub.contract.UnpackLog(event, "IntentSettled", log); err != nil {
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

// ParseIntentSettled is a log parse operation binding the contract event 0x3d59f7aa41c3b6ea14d68eb6d4fb841f9160a9b21ef5dca3b617d1d34eb712d4.
//
// Solidity: event IntentSettled(bytes32 indexed intentId)
func (_IntentHub *IntentHubFilterer) ParseIntentSettled(log types.Log) (*IntentHubIntentSettled, error) {
	event := new(IntentHubIntentSettled)
	if err := _IntentHub.contract.UnpackLog(event, "IntentSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
