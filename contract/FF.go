// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// FFMetaData contains all meta data concerning the FF contract.
var FFMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"}],\"name\":\"Borrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Lent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collateral\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakingBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lendingBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowingBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FFABI is the input ABI used to generate the binding from.
// Deprecated: Use FFMetaData.ABI instead.
var FFABI = FFMetaData.ABI

// FF is an auto generated Go binding around an Ethereum contract.
type FF struct {
	FFCaller     // Read-only binding to the contract
	FFTransactor // Write-only binding to the contract
	FFFilterer   // Log filterer for contract events
}

// FFCaller is an auto generated read-only Go binding around an Ethereum contract.
type FFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FFSession struct {
	Contract     *FF               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FFCallerSession struct {
	Contract *FFCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FFTransactorSession struct {
	Contract     *FFTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FFRaw is an auto generated low-level Go binding around an Ethereum contract.
type FFRaw struct {
	Contract *FF // Generic contract binding to access the raw methods on
}

// FFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FFCallerRaw struct {
	Contract *FFCaller // Generic read-only contract binding to access the raw methods on
}

// FFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FFTransactorRaw struct {
	Contract *FFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFF creates a new instance of FF, bound to a specific deployed contract.
func NewFF(address common.Address, backend bind.ContractBackend) (*FF, error) {
	contract, err := bindFF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FF{FFCaller: FFCaller{contract: contract}, FFTransactor: FFTransactor{contract: contract}, FFFilterer: FFFilterer{contract: contract}}, nil
}

// NewFFCaller creates a new read-only instance of FF, bound to a specific deployed contract.
func NewFFCaller(address common.Address, caller bind.ContractCaller) (*FFCaller, error) {
	contract, err := bindFF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FFCaller{contract: contract}, nil
}

// NewFFTransactor creates a new write-only instance of FF, bound to a specific deployed contract.
func NewFFTransactor(address common.Address, transactor bind.ContractTransactor) (*FFTransactor, error) {
	contract, err := bindFF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FFTransactor{contract: contract}, nil
}

// NewFFFilterer creates a new log filterer instance of FF, bound to a specific deployed contract.
func NewFFFilterer(address common.Address, filterer bind.ContractFilterer) (*FFFilterer, error) {
	contract, err := bindFF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FFFilterer{contract: contract}, nil
}

// bindFF binds a generic wrapper to an already deployed contract.
func bindFF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FF *FFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FF.Contract.FFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FF *FFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FF.Contract.FFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FF *FFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FF.Contract.FFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FF *FFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FF *FFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FF *FFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FF.Contract.contract.Transact(opts, method, params...)
}

// BorrowedAmount is a free data retrieval call binding the contract method 0x948adb8f.
//
// Solidity: function borrowedAmount(address ) view returns(uint256)
func (_FF *FFCaller) BorrowedAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FF.contract.Call(opts, &out, "borrowedAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowedAmount is a free data retrieval call binding the contract method 0x948adb8f.
//
// Solidity: function borrowedAmount(address ) view returns(uint256)
func (_FF *FFSession) BorrowedAmount(arg0 common.Address) (*big.Int, error) {
	return _FF.Contract.BorrowedAmount(&_FF.CallOpts, arg0)
}

// BorrowedAmount is a free data retrieval call binding the contract method 0x948adb8f.
//
// Solidity: function borrowedAmount(address ) view returns(uint256)
func (_FF *FFCallerSession) BorrowedAmount(arg0 common.Address) (*big.Int, error) {
	return _FF.Contract.BorrowedAmount(&_FF.CallOpts, arg0)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FF *FFCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FF.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FF *FFSession) Earned(account common.Address) (*big.Int, error) {
	return _FF.Contract.Earned(&_FF.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FF *FFCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _FF.Contract.Earned(&_FF.CallOpts, account)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_FF *FFCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FF.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_FF *FFSession) InterestRate() (*big.Int, error) {
	return _FF.Contract.InterestRate(&_FF.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_FF *FFCallerSession) InterestRate() (*big.Int, error) {
	return _FF.Contract.InterestRate(&_FF.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_FF *FFCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FF.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_FF *FFSession) RewardRate() (*big.Int, error) {
	return _FF.Contract.RewardRate(&_FF.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_FF *FFCallerSession) RewardRate() (*big.Int, error) {
	return _FF.Contract.RewardRate(&_FF.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 stakingBalance, uint256 lendingBalance, uint256 borrowingBalance, uint256 rewards, uint256 lastUpdated, uint256 collateralBalance)
func (_FF *FFCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	StakingBalance    *big.Int
	LendingBalance    *big.Int
	BorrowingBalance  *big.Int
	Rewards           *big.Int
	LastUpdated       *big.Int
	CollateralBalance *big.Int
}, error) {
	var out []interface{}
	err := _FF.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		StakingBalance    *big.Int
		LendingBalance    *big.Int
		BorrowingBalance  *big.Int
		Rewards           *big.Int
		LastUpdated       *big.Int
		CollateralBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakingBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LendingBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BorrowingBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdated = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CollateralBalance = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 stakingBalance, uint256 lendingBalance, uint256 borrowingBalance, uint256 rewards, uint256 lastUpdated, uint256 collateralBalance)
func (_FF *FFSession) Users(arg0 common.Address) (struct {
	StakingBalance    *big.Int
	LendingBalance    *big.Int
	BorrowingBalance  *big.Int
	Rewards           *big.Int
	LastUpdated       *big.Int
	CollateralBalance *big.Int
}, error) {
	return _FF.Contract.Users(&_FF.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(uint256 stakingBalance, uint256 lendingBalance, uint256 borrowingBalance, uint256 rewards, uint256 lastUpdated, uint256 collateralBalance)
func (_FF *FFCallerSession) Users(arg0 common.Address) (struct {
	StakingBalance    *big.Int
	LendingBalance    *big.Int
	BorrowingBalance  *big.Int
	Rewards           *big.Int
	LastUpdated       *big.Int
	CollateralBalance *big.Int
}, error) {
	return _FF.Contract.Users(&_FF.CallOpts, arg0)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 _amount, uint256 _collateral) returns()
func (_FF *FFTransactor) Borrow(opts *bind.TransactOpts, _amount *big.Int, _collateral *big.Int) (*types.Transaction, error) {
	return _FF.contract.Transact(opts, "borrow", _amount, _collateral)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 _amount, uint256 _collateral) returns()
func (_FF *FFSession) Borrow(_amount *big.Int, _collateral *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Borrow(&_FF.TransactOpts, _amount, _collateral)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 _amount, uint256 _collateral) returns()
func (_FF *FFTransactorSession) Borrow(_amount *big.Int, _collateral *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Borrow(&_FF.TransactOpts, _amount, _collateral)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_FF *FFTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FF.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_FF *FFSession) ClaimReward() (*types.Transaction, error) {
	return _FF.Contract.ClaimReward(&_FF.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_FF *FFTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _FF.Contract.ClaimReward(&_FF.TransactOpts)
}

// Lend is a paid mutator transaction binding the contract method 0xa6aa57ce.
//
// Solidity: function lend(uint256 _amount) returns()
func (_FF *FFTransactor) Lend(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FF.contract.Transact(opts, "lend", _amount)
}

// Lend is a paid mutator transaction binding the contract method 0xa6aa57ce.
//
// Solidity: function lend(uint256 _amount) returns()
func (_FF *FFSession) Lend(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Lend(&_FF.TransactOpts, _amount)
}

// Lend is a paid mutator transaction binding the contract method 0xa6aa57ce.
//
// Solidity: function lend(uint256 _amount) returns()
func (_FF *FFTransactorSession) Lend(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Lend(&_FF.TransactOpts, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _amount) returns()
func (_FF *FFTransactor) Repay(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FF.contract.Transact(opts, "repay", _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _amount) returns()
func (_FF *FFSession) Repay(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Repay(&_FF.TransactOpts, _amount)
}

// Repay is a paid mutator transaction binding the contract method 0x371fd8e6.
//
// Solidity: function repay(uint256 _amount) returns()
func (_FF *FFTransactorSession) Repay(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Repay(&_FF.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_FF *FFTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FF.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_FF *FFSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Stake(&_FF.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_FF *FFTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Stake(&_FF.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_FF *FFTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FF.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_FF *FFSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Withdraw(&_FF.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_FF *FFTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _FF.Contract.Withdraw(&_FF.TransactOpts, _amount)
}

// FFBorrowedIterator is returned from FilterBorrowed and is used to iterate over the raw logs and unpacked data for Borrowed events raised by the FF contract.
type FFBorrowedIterator struct {
	Event *FFBorrowed // Event containing the contract specifics and raw log

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
func (it *FFBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FFBorrowed)
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
		it.Event = new(FFBorrowed)
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
func (it *FFBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FFBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FFBorrowed represents a Borrowed event raised by the FF contract.
type FFBorrowed struct {
	User       common.Address
	Amount     *big.Int
	Collateral *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBorrowed is a free log retrieval operation binding the contract event 0xeae9cfbc77fdd40ca899f36b608256063b2bc9d8178b0220f7ad513e178d6730.
//
// Solidity: event Borrowed(address indexed user, uint256 amount, uint256 collateral)
func (_FF *FFFilterer) FilterBorrowed(opts *bind.FilterOpts, user []common.Address) (*FFBorrowedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.FilterLogs(opts, "Borrowed", userRule)
	if err != nil {
		return nil, err
	}
	return &FFBorrowedIterator{contract: _FF.contract, event: "Borrowed", logs: logs, sub: sub}, nil
}

// WatchBorrowed is a free log subscription operation binding the contract event 0xeae9cfbc77fdd40ca899f36b608256063b2bc9d8178b0220f7ad513e178d6730.
//
// Solidity: event Borrowed(address indexed user, uint256 amount, uint256 collateral)
func (_FF *FFFilterer) WatchBorrowed(opts *bind.WatchOpts, sink chan<- *FFBorrowed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.WatchLogs(opts, "Borrowed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FFBorrowed)
				if err := _FF.contract.UnpackLog(event, "Borrowed", log); err != nil {
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

// ParseBorrowed is a log parse operation binding the contract event 0xeae9cfbc77fdd40ca899f36b608256063b2bc9d8178b0220f7ad513e178d6730.
//
// Solidity: event Borrowed(address indexed user, uint256 amount, uint256 collateral)
func (_FF *FFFilterer) ParseBorrowed(log types.Log) (*FFBorrowed, error) {
	event := new(FFBorrowed)
	if err := _FF.contract.UnpackLog(event, "Borrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FFLentIterator is returned from FilterLent and is used to iterate over the raw logs and unpacked data for Lent events raised by the FF contract.
type FFLentIterator struct {
	Event *FFLent // Event containing the contract specifics and raw log

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
func (it *FFLentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FFLent)
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
		it.Event = new(FFLent)
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
func (it *FFLentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FFLentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FFLent represents a Lent event raised by the FF contract.
type FFLent struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLent is a free log retrieval operation binding the contract event 0x26ff3f6dedac6d273c293b8c866e073f40896739d99ad513862fe7ba70cf63ca.
//
// Solidity: event Lent(address indexed user, uint256 amount)
func (_FF *FFFilterer) FilterLent(opts *bind.FilterOpts, user []common.Address) (*FFLentIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.FilterLogs(opts, "Lent", userRule)
	if err != nil {
		return nil, err
	}
	return &FFLentIterator{contract: _FF.contract, event: "Lent", logs: logs, sub: sub}, nil
}

// WatchLent is a free log subscription operation binding the contract event 0x26ff3f6dedac6d273c293b8c866e073f40896739d99ad513862fe7ba70cf63ca.
//
// Solidity: event Lent(address indexed user, uint256 amount)
func (_FF *FFFilterer) WatchLent(opts *bind.WatchOpts, sink chan<- *FFLent, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.WatchLogs(opts, "Lent", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FFLent)
				if err := _FF.contract.UnpackLog(event, "Lent", log); err != nil {
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

// ParseLent is a log parse operation binding the contract event 0x26ff3f6dedac6d273c293b8c866e073f40896739d99ad513862fe7ba70cf63ca.
//
// Solidity: event Lent(address indexed user, uint256 amount)
func (_FF *FFFilterer) ParseLent(log types.Log) (*FFLent, error) {
	event := new(FFLent)
	if err := _FF.contract.UnpackLog(event, "Lent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FFRepaidIterator is returned from FilterRepaid and is used to iterate over the raw logs and unpacked data for Repaid events raised by the FF contract.
type FFRepaidIterator struct {
	Event *FFRepaid // Event containing the contract specifics and raw log

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
func (it *FFRepaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FFRepaid)
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
		it.Event = new(FFRepaid)
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
func (it *FFRepaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FFRepaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FFRepaid represents a Repaid event raised by the FF contract.
type FFRepaid struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRepaid is a free log retrieval operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_FF *FFFilterer) FilterRepaid(opts *bind.FilterOpts, user []common.Address) (*FFRepaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.FilterLogs(opts, "Repaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FFRepaidIterator{contract: _FF.contract, event: "Repaid", logs: logs, sub: sub}, nil
}

// WatchRepaid is a free log subscription operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_FF *FFFilterer) WatchRepaid(opts *bind.WatchOpts, sink chan<- *FFRepaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.WatchLogs(opts, "Repaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FFRepaid)
				if err := _FF.contract.UnpackLog(event, "Repaid", log); err != nil {
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

// ParseRepaid is a log parse operation binding the contract event 0x0516911bcc3a0a7412a44601057c0a0a1ec628bde049a84284bc428866534488.
//
// Solidity: event Repaid(address indexed user, uint256 amount)
func (_FF *FFFilterer) ParseRepaid(log types.Log) (*FFRepaid, error) {
	event := new(FFRepaid)
	if err := _FF.contract.UnpackLog(event, "Repaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FFRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the FF contract.
type FFRewardPaidIterator struct {
	Event *FFRewardPaid // Event containing the contract specifics and raw log

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
func (it *FFRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FFRewardPaid)
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
		it.Event = new(FFRewardPaid)
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
func (it *FFRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FFRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FFRewardPaid represents a RewardPaid event raised by the FF contract.
type FFRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_FF *FFFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*FFRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FFRewardPaidIterator{contract: _FF.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_FF *FFFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *FFRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FFRewardPaid)
				if err := _FF.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_FF *FFFilterer) ParseRewardPaid(log types.Log) (*FFRewardPaid, error) {
	event := new(FFRewardPaid)
	if err := _FF.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FFStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the FF contract.
type FFStakedIterator struct {
	Event *FFStaked // Event containing the contract specifics and raw log

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
func (it *FFStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FFStaked)
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
		it.Event = new(FFStaked)
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
func (it *FFStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FFStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FFStaked represents a Staked event raised by the FF contract.
type FFStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_FF *FFFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*FFStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &FFStakedIterator{contract: _FF.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_FF *FFFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *FFStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FFStaked)
				if err := _FF.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_FF *FFFilterer) ParseStaked(log types.Log) (*FFStaked, error) {
	event := new(FFStaked)
	if err := _FF.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FFWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the FF contract.
type FFWithdrawnIterator struct {
	Event *FFWithdrawn // Event containing the contract specifics and raw log

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
func (it *FFWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FFWithdrawn)
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
		it.Event = new(FFWithdrawn)
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
func (it *FFWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FFWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FFWithdrawn represents a Withdrawn event raised by the FF contract.
type FFWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_FF *FFFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*FFWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &FFWithdrawnIterator{contract: _FF.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_FF *FFFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *FFWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FF.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FFWithdrawn)
				if err := _FF.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_FF *FFFilterer) ParseWithdrawn(log types.Log) (*FFWithdrawn, error) {
	event := new(FFWithdrawn)
	if err := _FF.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
