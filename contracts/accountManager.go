// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AccountManagerABI is the input ABI used to generate the binding from.
const AccountManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"addAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes32\"}],\"name\":\"getAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"getUUID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashedRecoveryPhrase\",\"type\":\"bytes32\"}],\"name\":\"hasHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldRPHash\",\"type\":\"bytes32\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"newRPHash\",\"type\":\"bytes32\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getAssociatedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes32\"},{\"name\":\"hashedRecoveryPhrase\",\"type\":\"bytes32\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAuthKeyAddress\",\"type\":\"address\"}],\"name\":\"changeAuthKeyAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"newAuthKeyAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newKeyAddress\",\"type\":\"address\"}],\"name\":\"AuthorisationKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AccountRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes32\"}],\"name\":\"AccountsAdded\",\"type\":\"event\"}]"

// AccountManagerBin is the compiled bytecode used for deploying new contracts.
const AccountManagerBin = `0x608060405234801561001057600080fd5b50604051602080610993833981016040525160008054600160a060020a03338116600160a060020a031992831617909255600180549290931691161790556109368061005d6000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633628731c81146100b35780633fdada7a1461010a5780638da5cb5b1461017257806392da867f146101a35780639f5baacc146101e2578063a8a043991461020e578063c3c5a54714610235578063c7c04cc014610256578063eea9542914610277578063f2fde38b1461029e578063ff51886f146102bf575b600080fd5b3480156100bf57600080fd5b5060408051602060048035808201358381028086018501909652808552610108953695939460249493850192918291850190849080828437509497506102e09650505050505050565b005b34801561011657600080fd5b506101226004356103c3565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561015e578181015183820152602001610146565b505050509050019250505060405180910390f35b34801561017e57600080fd5b50610187610443565b60408051600160a060020a039092168252519081900360200190f35b3480156101af57600080fd5b506101d0600160a060020a036004351660243560443560ff60643516610452565b60408051918252519081900360200190f35b3480156101ee57600080fd5b506101fa6004356104cb565b604080519115158252519081900360200190f35b34801561021a57600080fd5b506101fa60043560243560443560ff606435166084356104f5565b34801561024157600080fd5b506101fa600160a060020a03600435166105c4565b34801561026257600080fd5b50610122600160a060020a03600435166105da565b34801561028357600080fd5b506101fa60043560243560443560643560ff60843516610674565b3480156102aa57600080fd5b50610108600160a060020a036004351661078b565b3480156102cb57600080fd5b50610108600160a060020a0360043516610807565b6000806102ec336105c4565b15156102f757600080fd5b5050600160a060020a033316600090815260036020526040812054905b825181101561038b576000828152600260205260409020835184908390811061033957fe5b602090810291909101810151825460018082018555600094855292909320909201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039093169290921790915501610314565b6040805183815290517fc1d0d899ac6587729c022ff99d15f711f7b8df0a288603edcbe1dd2b59951df89181900360200190a1505050565b60606103ce336105c4565b15156103d957600080fd5b6000828152600260209081526040918290208054835181840281018401909452808452909183018282801561043757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610419575b50505050509050919050565b600054600160a060020a031681565b6000805433600160a060020a0390811691161461046e57600080fd5b604080516c01000000000000000000000000600160a060020a03881602815290519081900360140190206104a490858585610883565b15156104ac57fe5b50505050600160a060020a031660009081526003602052604090205490565b6000805433600160a060020a039081169116146104e757600080fd5b506000526004602052600190565b600085815260046020528061050c87878787610883565b151561051457fe5b5060008681526004602081815260408084208054878652828620819055600160a060020a0333168087526003855283872082905581875260028552838720805460018101825590885285882001805473ffffffffffffffffffffffffffffffffffffffff1916821790558c87529484529490558051928352517f444b9fc80713a03bfc01b03089babd94aa51edececce36355031fd7d1ed50b4a9281900390910190a15060019695505050505050565b600160a060020a03166000908152600360205290565b60005460609033600160a060020a039081169116146105f857600080fd5b600160a060020a038216600090815260036020908152604080832054835260028252918290208054835181840281018401909452808452909183018282801561043757602002820191906000526020600020908154600160a060020a031681526001909101906020018083116104195750505050509050919050565b6000805433600160a060020a039081169116141561069157600080fd5b61069a336105c4565b156106a457600080fd5b604080516c01000000000000000000000000600160a060020a03331602815290519081900360140190206106da90858585610883565b15156106e557600080fd5b600160a060020a03331660008181526003602090815260408083208a905589835260028252808320805460018101825590845282842001805473ffffffffffffffffffffffffffffffffffffffff191685179055888352600482529182902089905581518981529081019290925280517f7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c9281900390910190a150600195945050505050565b60005433600160a060020a039081169116146107a657600080fd5b60008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc6861639181900360200190a150565b60005433600160a060020a0390811691161461082257600080fd5b60018054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252517f397e9e5bff25b6d6a7ba18f106ec27385780db7fc612fc66924903e8ef2144f49181900360200190a150565b6001805460408051600080825260208083018085528a905260ff871683850152606083018990526080830188905292519094600160a060020a03909416939260a0808401939192601f198101928190039091019087865af11580156108ec573d6000803e3d6000fd5b50505060206040510351600160a060020a03161490509493505050505600a165627a7a72305820f3db3974bf03f8f319c6e73228d55a4a89616c9d568ac1d1185008b6a6ef2f9f0029`

// DeployAccountManager deploys a new Ethereum contract, binding an instance of AccountManager to it.
func DeployAccountManager(auth *bind.TransactOpts, backend bind.ContractBackend, newAuthKeyAddress common.Address) (common.Address, *types.Transaction, *AccountManager, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountManagerBin), backend, newAuthKeyAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// AccountManager is an auto generated Go binding around an Ethereum contract.
type AccountManager struct {
	AccountManagerCaller     // Read-only binding to the contract
	AccountManagerTransactor // Write-only binding to the contract
	AccountManagerFilterer   // Log filterer for contract events
}

// AccountManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerSession struct {
	Contract     *AccountManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerCallerSession struct {
	Contract *AccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerTransactorSession struct {
	Contract     *AccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerRaw struct {
	Contract *AccountManager // Generic contract binding to access the raw methods on
}

// AccountManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerCallerRaw struct {
	Contract *AccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerTransactorRaw struct {
	Contract *AccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManager creates a new instance of AccountManager, bound to a specific deployed contract.
func NewAccountManager(address common.Address, backend bind.ContractBackend) (*AccountManager, error) {
	contract, err := bindAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// NewAccountManagerCaller creates a new read-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerCaller, error) {
	contract, err := bindAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCaller{contract: contract}, nil
}

// NewAccountManagerTransactor creates a new write-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerTransactor, error) {
	contract, err := bindAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerTransactor{contract: contract}, nil
}

// NewAccountManagerFilterer creates a new log filterer instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFilterer, error) {
	contract, err := bindAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFilterer{contract: contract}, nil
}

// bindAccountManager binds a generic wrapper to an already deployed contract.
func bindAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.AccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transact(opts, method, params...)
}

// GetAddresses is a free data retrieval call binding the contract method 0x3fdada7a.
//
// Solidity: function getAddresses(uuid bytes32) constant returns(address[])
func (_AccountManager *AccountManagerCaller) GetAddresses(opts *bind.CallOpts, uuid [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getAddresses", uuid)
	return *ret0, err
}

// GetAddresses is a free data retrieval call binding the contract method 0x3fdada7a.
//
// Solidity: function getAddresses(uuid bytes32) constant returns(address[])
func (_AccountManager *AccountManagerSession) GetAddresses(uuid [32]byte) ([]common.Address, error) {
	return _AccountManager.Contract.GetAddresses(&_AccountManager.CallOpts, uuid)
}

// GetAddresses is a free data retrieval call binding the contract method 0x3fdada7a.
//
// Solidity: function getAddresses(uuid bytes32) constant returns(address[])
func (_AccountManager *AccountManagerCallerSession) GetAddresses(uuid [32]byte) ([]common.Address, error) {
	return _AccountManager.Contract.GetAddresses(&_AccountManager.CallOpts, uuid)
}

// GetAssociatedAddresses is a free data retrieval call binding the contract method 0xc7c04cc0.
//
// Solidity: function getAssociatedAddresses(userAddress address) constant returns(address[])
func (_AccountManager *AccountManagerCaller) GetAssociatedAddresses(opts *bind.CallOpts, userAddress common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getAssociatedAddresses", userAddress)
	return *ret0, err
}

// GetAssociatedAddresses is a free data retrieval call binding the contract method 0xc7c04cc0.
//
// Solidity: function getAssociatedAddresses(userAddress address) constant returns(address[])
func (_AccountManager *AccountManagerSession) GetAssociatedAddresses(userAddress common.Address) ([]common.Address, error) {
	return _AccountManager.Contract.GetAssociatedAddresses(&_AccountManager.CallOpts, userAddress)
}

// GetAssociatedAddresses is a free data retrieval call binding the contract method 0xc7c04cc0.
//
// Solidity: function getAssociatedAddresses(userAddress address) constant returns(address[])
func (_AccountManager *AccountManagerCallerSession) GetAssociatedAddresses(userAddress common.Address) ([]common.Address, error) {
	return _AccountManager.Contract.GetAssociatedAddresses(&_AccountManager.CallOpts, userAddress)
}

// GetUUID is a free data retrieval call binding the contract method 0x92da867f.
//
// Solidity: function getUUID(userAddress address, sigR bytes32, sigS bytes32, sigV uint8) constant returns(bytes32)
func (_AccountManager *AccountManagerCaller) GetUUID(opts *bind.CallOpts, userAddress common.Address, sigR [32]byte, sigS [32]byte, sigV uint8) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getUUID", userAddress, sigR, sigS, sigV)
	return *ret0, err
}

// GetUUID is a free data retrieval call binding the contract method 0x92da867f.
//
// Solidity: function getUUID(userAddress address, sigR bytes32, sigS bytes32, sigV uint8) constant returns(bytes32)
func (_AccountManager *AccountManagerSession) GetUUID(userAddress common.Address, sigR [32]byte, sigS [32]byte, sigV uint8) ([32]byte, error) {
	return _AccountManager.Contract.GetUUID(&_AccountManager.CallOpts, userAddress, sigR, sigS, sigV)
}

// GetUUID is a free data retrieval call binding the contract method 0x92da867f.
//
// Solidity: function getUUID(userAddress address, sigR bytes32, sigS bytes32, sigV uint8) constant returns(bytes32)
func (_AccountManager *AccountManagerCallerSession) GetUUID(userAddress common.Address, sigR [32]byte, sigS [32]byte, sigV uint8) ([32]byte, error) {
	return _AccountManager.Contract.GetUUID(&_AccountManager.CallOpts, userAddress, sigR, sigS, sigV)
}

// HasHash is a free data retrieval call binding the contract method 0x9f5baacc.
//
// Solidity: function hasHash(hashedRecoveryPhrase bytes32) constant returns(bool)
func (_AccountManager *AccountManagerCaller) HasHash(opts *bind.CallOpts, hashedRecoveryPhrase [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "hasHash", hashedRecoveryPhrase)
	return *ret0, err
}

// HasHash is a free data retrieval call binding the contract method 0x9f5baacc.
//
// Solidity: function hasHash(hashedRecoveryPhrase bytes32) constant returns(bool)
func (_AccountManager *AccountManagerSession) HasHash(hashedRecoveryPhrase [32]byte) (bool, error) {
	return _AccountManager.Contract.HasHash(&_AccountManager.CallOpts, hashedRecoveryPhrase)
}

// HasHash is a free data retrieval call binding the contract method 0x9f5baacc.
//
// Solidity: function hasHash(hashedRecoveryPhrase bytes32) constant returns(bool)
func (_AccountManager *AccountManagerCallerSession) HasHash(hashedRecoveryPhrase [32]byte) (bool, error) {
	return _AccountManager.Contract.HasHash(&_AccountManager.CallOpts, hashedRecoveryPhrase)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(user address) constant returns(bool)
func (_AccountManager *AccountManagerCaller) IsRegistered(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "isRegistered", user)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(user address) constant returns(bool)
func (_AccountManager *AccountManagerSession) IsRegistered(user common.Address) (bool, error) {
	return _AccountManager.Contract.IsRegistered(&_AccountManager.CallOpts, user)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(user address) constant returns(bool)
func (_AccountManager *AccountManagerCallerSession) IsRegistered(user common.Address) (bool, error) {
	return _AccountManager.Contract.IsRegistered(&_AccountManager.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountManager *AccountManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountManager *AccountManagerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountManager *AccountManagerCallerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// AddAddresses is a paid mutator transaction binding the contract method 0x3628731c.
//
// Solidity: function addAddresses(addresses address[]) returns()
func (_AccountManager *AccountManagerTransactor) AddAddresses(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "addAddresses", addresses)
}

// AddAddresses is a paid mutator transaction binding the contract method 0x3628731c.
//
// Solidity: function addAddresses(addresses address[]) returns()
func (_AccountManager *AccountManagerSession) AddAddresses(addresses []common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AddAddresses(&_AccountManager.TransactOpts, addresses)
}

// AddAddresses is a paid mutator transaction binding the contract method 0x3628731c.
//
// Solidity: function addAddresses(addresses address[]) returns()
func (_AccountManager *AccountManagerTransactorSession) AddAddresses(addresses []common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AddAddresses(&_AccountManager.TransactOpts, addresses)
}

// ChangeAuthKeyAddress is a paid mutator transaction binding the contract method 0xff51886f.
//
// Solidity: function changeAuthKeyAddress(newAuthKeyAddress address) returns()
func (_AccountManager *AccountManagerTransactor) ChangeAuthKeyAddress(opts *bind.TransactOpts, newAuthKeyAddress common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "changeAuthKeyAddress", newAuthKeyAddress)
}

// ChangeAuthKeyAddress is a paid mutator transaction binding the contract method 0xff51886f.
//
// Solidity: function changeAuthKeyAddress(newAuthKeyAddress address) returns()
func (_AccountManager *AccountManagerSession) ChangeAuthKeyAddress(newAuthKeyAddress common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.ChangeAuthKeyAddress(&_AccountManager.TransactOpts, newAuthKeyAddress)
}

// ChangeAuthKeyAddress is a paid mutator transaction binding the contract method 0xff51886f.
//
// Solidity: function changeAuthKeyAddress(newAuthKeyAddress address) returns()
func (_AccountManager *AccountManagerTransactorSession) ChangeAuthKeyAddress(newAuthKeyAddress common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.ChangeAuthKeyAddress(&_AccountManager.TransactOpts, newAuthKeyAddress)
}

// Recover is a paid mutator transaction binding the contract method 0xa8a04399.
//
// Solidity: function recover(oldRPHash bytes32, sigR bytes32, sigS bytes32, sigV uint8, newRPHash bytes32) returns(bool)
func (_AccountManager *AccountManagerTransactor) Recover(opts *bind.TransactOpts, oldRPHash [32]byte, sigR [32]byte, sigS [32]byte, sigV uint8, newRPHash [32]byte) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "recover", oldRPHash, sigR, sigS, sigV, newRPHash)
}

// Recover is a paid mutator transaction binding the contract method 0xa8a04399.
//
// Solidity: function recover(oldRPHash bytes32, sigR bytes32, sigS bytes32, sigV uint8, newRPHash bytes32) returns(bool)
func (_AccountManager *AccountManagerSession) Recover(oldRPHash [32]byte, sigR [32]byte, sigS [32]byte, sigV uint8, newRPHash [32]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Recover(&_AccountManager.TransactOpts, oldRPHash, sigR, sigS, sigV, newRPHash)
}

// Recover is a paid mutator transaction binding the contract method 0xa8a04399.
//
// Solidity: function recover(oldRPHash bytes32, sigR bytes32, sigS bytes32, sigV uint8, newRPHash bytes32) returns(bool)
func (_AccountManager *AccountManagerTransactorSession) Recover(oldRPHash [32]byte, sigR [32]byte, sigS [32]byte, sigV uint8, newRPHash [32]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Recover(&_AccountManager.TransactOpts, oldRPHash, sigR, sigS, sigV, newRPHash)
}

// Register is a paid mutator transaction binding the contract method 0xeea95429.
//
// Solidity: function register(uuid bytes32, hashedRecoveryPhrase bytes32, sigR bytes32, sigS bytes32, sigV uint8) returns(bool)
func (_AccountManager *AccountManagerTransactor) Register(opts *bind.TransactOpts, uuid [32]byte, hashedRecoveryPhrase [32]byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "register", uuid, hashedRecoveryPhrase, sigR, sigS, sigV)
}

// Register is a paid mutator transaction binding the contract method 0xeea95429.
//
// Solidity: function register(uuid bytes32, hashedRecoveryPhrase bytes32, sigR bytes32, sigS bytes32, sigV uint8) returns(bool)
func (_AccountManager *AccountManagerSession) Register(uuid [32]byte, hashedRecoveryPhrase [32]byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _AccountManager.Contract.Register(&_AccountManager.TransactOpts, uuid, hashedRecoveryPhrase, sigR, sigS, sigV)
}

// Register is a paid mutator transaction binding the contract method 0xeea95429.
//
// Solidity: function register(uuid bytes32, hashedRecoveryPhrase bytes32, sigR bytes32, sigS bytes32, sigV uint8) returns(bool)
func (_AccountManager *AccountManagerTransactorSession) Register(uuid [32]byte, hashedRecoveryPhrase [32]byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _AccountManager.Contract.Register(&_AccountManager.TransactOpts, uuid, hashedRecoveryPhrase, sigR, sigS, sigV)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AccountManager *AccountManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AccountManager *AccountManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.TransferOwnership(&_AccountManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AccountManager *AccountManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.TransferOwnership(&_AccountManager.TransactOpts, newOwner)
}

// AccountManagerAccountRecoveredIterator is returned from FilterAccountRecovered and is used to iterate over the raw logs and unpacked data for AccountRecovered events raised by the AccountManager contract.
type AccountManagerAccountRecoveredIterator struct {
	Event *AccountManagerAccountRecovered // Event containing the contract specifics and raw log

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
func (it *AccountManagerAccountRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAccountRecovered)
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
		it.Event = new(AccountManagerAccountRecovered)
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
func (it *AccountManagerAccountRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAccountRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAccountRecovered represents a AccountRecovered event raised by the AccountManager contract.
type AccountManagerAccountRecovered struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountRecovered is a free log retrieval operation binding the contract event 0x444b9fc80713a03bfc01b03089babd94aa51edececce36355031fd7d1ed50b4a.
//
// Solidity: event AccountRecovered(newAddress address)
func (_AccountManager *AccountManagerFilterer) FilterAccountRecovered(opts *bind.FilterOpts) (*AccountManagerAccountRecoveredIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAccountRecoveredIterator{contract: _AccountManager.contract, event: "AccountRecovered", logs: logs, sub: sub}, nil
}

// WatchAccountRecovered is a free log subscription operation binding the contract event 0x444b9fc80713a03bfc01b03089babd94aa51edececce36355031fd7d1ed50b4a.
//
// Solidity: event AccountRecovered(newAddress address)
func (_AccountManager *AccountManagerFilterer) WatchAccountRecovered(opts *bind.WatchOpts, sink chan<- *AccountManagerAccountRecovered) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAccountRecovered)
				if err := _AccountManager.contract.UnpackLog(event, "AccountRecovered", log); err != nil {
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

// AccountManagerAccountsAddedIterator is returned from FilterAccountsAdded and is used to iterate over the raw logs and unpacked data for AccountsAdded events raised by the AccountManager contract.
type AccountManagerAccountsAddedIterator struct {
	Event *AccountManagerAccountsAdded // Event containing the contract specifics and raw log

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
func (it *AccountManagerAccountsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAccountsAdded)
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
		it.Event = new(AccountManagerAccountsAdded)
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
func (it *AccountManagerAccountsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAccountsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAccountsAdded represents a AccountsAdded event raised by the AccountManager contract.
type AccountManagerAccountsAdded struct {
	Uuid [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAccountsAdded is a free log retrieval operation binding the contract event 0xc1d0d899ac6587729c022ff99d15f711f7b8df0a288603edcbe1dd2b59951df8.
//
// Solidity: event AccountsAdded(uuid bytes32)
func (_AccountManager *AccountManagerFilterer) FilterAccountsAdded(opts *bind.FilterOpts) (*AccountManagerAccountsAddedIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "AccountsAdded")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAccountsAddedIterator{contract: _AccountManager.contract, event: "AccountsAdded", logs: logs, sub: sub}, nil
}

// WatchAccountsAdded is a free log subscription operation binding the contract event 0xc1d0d899ac6587729c022ff99d15f711f7b8df0a288603edcbe1dd2b59951df8.
//
// Solidity: event AccountsAdded(uuid bytes32)
func (_AccountManager *AccountManagerFilterer) WatchAccountsAdded(opts *bind.WatchOpts, sink chan<- *AccountManagerAccountsAdded) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "AccountsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAccountsAdded)
				if err := _AccountManager.contract.UnpackLog(event, "AccountsAdded", log); err != nil {
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

// AccountManagerAuthorisationKeyUpdatedIterator is returned from FilterAuthorisationKeyUpdated and is used to iterate over the raw logs and unpacked data for AuthorisationKeyUpdated events raised by the AccountManager contract.
type AccountManagerAuthorisationKeyUpdatedIterator struct {
	Event *AccountManagerAuthorisationKeyUpdated // Event containing the contract specifics and raw log

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
func (it *AccountManagerAuthorisationKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerAuthorisationKeyUpdated)
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
		it.Event = new(AccountManagerAuthorisationKeyUpdated)
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
func (it *AccountManagerAuthorisationKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerAuthorisationKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerAuthorisationKeyUpdated represents a AuthorisationKeyUpdated event raised by the AccountManager contract.
type AccountManagerAuthorisationKeyUpdated struct {
	NewKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorisationKeyUpdated is a free log retrieval operation binding the contract event 0x397e9e5bff25b6d6a7ba18f106ec27385780db7fc612fc66924903e8ef2144f4.
//
// Solidity: event AuthorisationKeyUpdated(newKeyAddress address)
func (_AccountManager *AccountManagerFilterer) FilterAuthorisationKeyUpdated(opts *bind.FilterOpts) (*AccountManagerAuthorisationKeyUpdatedIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "AuthorisationKeyUpdated")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAuthorisationKeyUpdatedIterator{contract: _AccountManager.contract, event: "AuthorisationKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorisationKeyUpdated is a free log subscription operation binding the contract event 0x397e9e5bff25b6d6a7ba18f106ec27385780db7fc612fc66924903e8ef2144f4.
//
// Solidity: event AuthorisationKeyUpdated(newKeyAddress address)
func (_AccountManager *AccountManagerFilterer) WatchAuthorisationKeyUpdated(opts *bind.WatchOpts, sink chan<- *AccountManagerAuthorisationKeyUpdated) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "AuthorisationKeyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerAuthorisationKeyUpdated)
				if err := _AccountManager.contract.UnpackLog(event, "AuthorisationKeyUpdated", log); err != nil {
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

// AccountManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccountManager contract.
type AccountManagerOwnershipTransferredIterator struct {
	Event *AccountManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerOwnershipTransferred)
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
		it.Event = new(AccountManagerOwnershipTransferred)
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
func (it *AccountManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AccountManager contract.
type AccountManagerOwnershipTransferred struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(newOwner address)
func (_AccountManager *AccountManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts) (*AccountManagerOwnershipTransferredIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return &AccountManagerOwnershipTransferredIterator{contract: _AccountManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc686163.
//
// Solidity: event OwnershipTransferred(newOwner address)
func (_AccountManager *AccountManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountManagerOwnershipTransferred) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "OwnershipTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerOwnershipTransferred)
				if err := _AccountManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AccountManagerRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the AccountManager contract.
type AccountManagerRegisteredIterator struct {
	Event *AccountManagerRegistered // Event containing the contract specifics and raw log

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
func (it *AccountManagerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerRegistered)
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
		it.Event = new(AccountManagerRegistered)
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
func (it *AccountManagerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerRegistered represents a Registered event raised by the AccountManager contract.
type AccountManagerRegistered struct {
	Uuid   [32]byte
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c.
//
// Solidity: event Registered(uuid bytes32, sender address)
func (_AccountManager *AccountManagerFilterer) FilterRegistered(opts *bind.FilterOpts) (*AccountManagerRegisteredIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &AccountManagerRegisteredIterator{contract: _AccountManager.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c.
//
// Solidity: event Registered(uuid bytes32, sender address)
func (_AccountManager *AccountManagerFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *AccountManagerRegistered) (event.Subscription, error) {

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerRegistered)
				if err := _AccountManager.contract.UnpackLog(event, "Registered", log); err != nil {
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

