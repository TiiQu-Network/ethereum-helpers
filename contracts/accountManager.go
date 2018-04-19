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
const AccountManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes32\"},{\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"addAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registrationKey\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes32\"},{\"name\":\"signedPublicKey\",\"type\":\"bytes32\"},{\"name\":\"hashedRecoveryPhrase\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashedMessage\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oldRPHash\",\"type\":\"bytes32\"},{\"name\":\"oldSignedRPHash\",\"type\":\"bytes32\"},{\"name\":\"newRPHash\",\"type\":\"bytes32\"},{\"name\":\"newSignedRPHash\",\"type\":\"bytes32\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userPublicKey\",\"type\":\"address\"}],\"name\":\"getUUID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"bytes32\"}],\"name\":\"getAdresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashedRecoveryPhrase\",\"type\":\"bytes32\"}],\"name\":\"hasHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userPublicKey\",\"type\":\"address\"}],\"name\":\"getAssociatedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userPublicKey\",\"type\":\"address\"}],\"name\":\"getUUIDString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_registrationKey\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes32\"}],\"name\":\"AccountRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"bytes32\"}],\"name\":\"AccountsAdded\",\"type\":\"event\"}]"

// AccountManagerBin is the compiled bytecode used for deploying new contracts.
const AccountManagerBin = `0x6060604052341561000f57600080fd5b6040516020806110bf8339810160405280805160008054600160a060020a03338116600160a060020a0319928316179092556001805492909316911617905550506110608061005f6000396000f3006060604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166336de7c1181146100c95780633b1978281461011f5780634da274fd1461014e57806375e366161461017e5780637f52ebb2146101945780638c01069f146101b35780638da5cb5b146101e45780639a55f103146101f75780639f5baacc14610260578063c3c5a54714610276578063c7c04cc014610295578063d669582c146102b4578063f2fde38b1461034a575b600080fd5b34156100d457600080fd5b61011d6004803590604460248035908101908301358060208181020160405190810160405280939291908181526020018383602002808284375094965061036995505050505050565b005b341561012a57600080fd5b610132610430565b604051600160a060020a03909116815260200160405180910390f35b341561015957600080fd5b61016a60043560243560443561043f565b604051901515815260200160405180910390f35b341561018957600080fd5b61016a600435610532565b341561019f57600080fd5b61016a6004356024356044356064356105cb565b34156101be57600080fd5b6101d2600160a060020a0360043516610709565b60405190815260200160405180910390f35b34156101ef57600080fd5b61013261073b565b341561020257600080fd5b61020d60043561074a565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561024c578082015183820152602001610234565b505050509050019250505060405180910390f35b341561026b57600080fd5b61016a6004356107d8565b341561028157600080fd5b61016a600160a060020a0360043516610802565b34156102a057600080fd5b61020d600160a060020a0360043516610818565b34156102bf57600080fd5b6102d3600160a060020a03600435166108bc565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561030f5780820151838201526020016102f7565b50505050905090810190601f16801561033c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561035557600080fd5b61011d600160a060020a0360043516610906565b600061037433610802565b151561037f57600080fd5b5060005b81518110156103f85760008381526002602052604090208054600181016103aa8382610fd8565b916000526020600020900160008484815181106103c357fe5b906020019060200201518254600160a060020a039182166101009390930a928302919092021990911617905550600101610383565b7fc1d0d899ac6587729c022ff99d15f711f7b8df0a288603edcbe1dd2b59951df88360405190815260200160405180910390a1505050565b600154600160a060020a031681565b600061044a33610802565b1561045157fe5b61045a83610532565b151561046257fe5b600160a060020a03331660009081526003602090815260408083208790558683526002909152902080546001810161049a8382610fd8565b506000918252602080832091909101805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a0381169190911790915584835260049091526040918290208690557f7d917fcbc9a29a9705ff9936ffa599500e4fd902e4486bae317414fe967b307c9186919051918252600160a060020a031660208201526040908101905180910390a15060019392505050565b60008060008061054185610989565b600180549396509194509250600160a060020a0390911690868585856040516000815260200160405260405193845260ff9092166020808501919091526040808501929092526060840192909252608090920191516020810390808403906000865af115156105af57600080fd5b505060206040510351600160a060020a03161495945050505050565b6000806105d733610802565b15156105e257600080fd5b600086905260046020526105f586610532565b15156105fd57fe5b61060684610532565b151561060e57fe5b61061785610532565b151561061f57fe5b61062883610532565b151561063057fe5b50600085815260046020908152604080832054868452818420819055600160a060020a033316845260038352818420819055808452600290925282209091906106799082610fd8565b5060008181526002602052604090208054600181016106988382610fd8565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff191633600160a060020a03161790557fd67568f31e693c2a76b51fa11a1edf173364e7f7c90f3bf34f802b4a49b8344d8160405190815260200160405180910390a150600195945050505050565b600061071433610802565b151561071f57600080fd5b50600160a060020a031660009081526003602052604090205490565b600054600160a060020a031681565b610752611001565b61075b33610802565b151561076657600080fd5b6000828152600260209081526040918290208054909290918281020190519081016040528092919081815260200182805480156107cc57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107ae575b50505050509050919050565b6000805433600160a060020a039081169116146107f457600080fd5b506000526004602052600190565b600160a060020a03166000908152600360205290565b610820611001565b60005433600160a060020a0390811691161461083b57600080fd5b6002600061084884610709565b600019166000191681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156107cc57602002820191906000526020600020908154600160a060020a031681526001909101906020018083116107ae5750505050509050919050565b6108c4611001565b60006108cf33610802565b15156108da57600080fd5b50600160a060020a0382166000908152600360205260409020546108fd81610a2a565b91505b50919050565b60005433600160a060020a0390811691161461092157600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f04dba622d284ed0014ee4b9a6a68386be1a4c08a4913ae272de89199cc68616381604051600160a060020a03909116815260200160405180910390a150565b60008060008060008061099a611001565b60005433600160a060020a039081169116146109b557600080fd5b6109d26109cd6109c48a610a2a565b60026084610b72565b610c41565b90506020810151925060408101519150606081015160001a9350601b8460ff1610156109ff57601b840193505b601b8460ff161180610a145750601c8460ff16105b1515610a1c57fe5b509196909550909350915050565b610a32611001565b610a3a611001565b6000806000610a47611001565b6020604051805910610a565750595b818152601f19601f83011681016020016040529050945060009350600092505b6020831015610ae9576008830260020a870291507fff00000000000000000000000000000000000000000000000000000000000000821615610ade5781858581518110610abf57fe5b906020010190600160f860020a031916908160001a9053506001909301925b600190920191610a76565b83604051805910610af75750595b818152601f19601f830116810160200160405290509050600092505b83831015610b6857848381518110610b2757fe5b016020015160f860020a900460f860020a02818481518110610b4557fe5b906020010190600160f860020a031916908160001a905350600190920191610b13565b9695505050505050565b610b7a611001565b610b82611001565b610b8a611001565b859150600084861115610b9c57600080fd5b6000861015610baa57600080fd5b8251851115610bb857600080fd5b858503604051805910610bc85750595b818152601f19601f8301168101602001604052905091508590505b84811015610c3757828181518110610bf757fe5b016020015160f860020a900460f860020a028287830381518110610c1757fe5b906020010190600160f860020a031916908160001a905350600101610be3565b5095945050505050565b610c49611001565b6000610c53611001565b6000610c5d611001565b610c65611001565b600080885196506002870615610c7a57600080fd5b60028704604051805910610c8b5750595b818152601f19601f83011681016020016040529050955060009450600091505b86821015610d4a57610cc1898384600101610b72565b9350610cd4898360010184600201610b72565b9250610cdf83610d57565b610ce885610d57565b601002019050610cf781610fa0565b601f81518110610d0357fe5b016020015160f860020a900460f860020a02868680600101975081518110610d2757fe5b906020010190600160f860020a031916908160001a905350600282019150610cab565b5093979650505050505050565b6000610d61611001565b50817f300000000000000000000000000000000000000000000000000000000000000081600081518110610d9157fe5b016020015160f860020a900460f860020a02600160f860020a03191610158015610e0357507f390000000000000000000000000000000000000000000000000000000000000081600081518110610de457fe5b016020015160f860020a900460f860020a02600160f860020a03191611155b15610e3857603081600081518110610e1757fe5b016020015160f860020a900460f860020a0260f860020a9004039150610900565b7f410000000000000000000000000000000000000000000000000000000000000081600081518110610e6657fe5b016020015160f860020a900460f860020a02600160f860020a03191610158015610ed857507f460000000000000000000000000000000000000000000000000000000000000081600081518110610eb957fe5b016020015160f860020a900460f860020a02600160f860020a03191611155b15610eec57603781600081518110610e1757fe5b7f610000000000000000000000000000000000000000000000000000000000000081600081518110610f1a57fe5b016020015160f860020a900460f860020a02600160f860020a03191610158015610f8c57507f660000000000000000000000000000000000000000000000000000000000000081600081518110610f6d57fe5b016020015160f860020a900460f860020a02600160f860020a03191611155b156100c457605781600081518110610e1757fe5b610fa8611001565b6020604051805910610fb75750595b818152601f19601f8301168101602001604052905060208101929092525090565b815481835581811511610ffc57600083815260209020610ffc918101908301611013565b505050565b60206040519081016040526000815290565b61103191905b8082111561102d5760008155600101611019565b5090565b905600a165627a7a72305820f6cebf3affb5dbfc19d29a0eeb2ae5a04b79c916282eebe0578c4fc0f091a2990029`

// DeployAccountManager deploys a new Ethereum contract, binding an instance of AccountManager to it.
func DeployAccountManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registrationKey common.Address) (common.Address, *types.Transaction, *AccountManager, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountManagerBin), backend, _registrationKey)
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

// GetAdresses is a free data retrieval call binding the contract method 0x9a55f103.
//
// Solidity: function getAdresses(uuid bytes32) constant returns(address[])
func (_AccountManager *AccountManagerCaller) GetAdresses(opts *bind.CallOpts, uuid [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getAdresses", uuid)
	return *ret0, err
}

// GetAdresses is a free data retrieval call binding the contract method 0x9a55f103.
//
// Solidity: function getAdresses(uuid bytes32) constant returns(address[])
func (_AccountManager *AccountManagerSession) GetAdresses(uuid [32]byte) ([]common.Address, error) {
	return _AccountManager.Contract.GetAdresses(&_AccountManager.CallOpts, uuid)
}

// GetAdresses is a free data retrieval call binding the contract method 0x9a55f103.
//
// Solidity: function getAdresses(uuid bytes32) constant returns(address[])
func (_AccountManager *AccountManagerCallerSession) GetAdresses(uuid [32]byte) ([]common.Address, error) {
	return _AccountManager.Contract.GetAdresses(&_AccountManager.CallOpts, uuid)
}

// GetAssociatedAddresses is a free data retrieval call binding the contract method 0xc7c04cc0.
//
// Solidity: function getAssociatedAddresses(userPublicKey address) constant returns(address[])
func (_AccountManager *AccountManagerCaller) GetAssociatedAddresses(opts *bind.CallOpts, userPublicKey common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getAssociatedAddresses", userPublicKey)
	return *ret0, err
}

// GetAssociatedAddresses is a free data retrieval call binding the contract method 0xc7c04cc0.
//
// Solidity: function getAssociatedAddresses(userPublicKey address) constant returns(address[])
func (_AccountManager *AccountManagerSession) GetAssociatedAddresses(userPublicKey common.Address) ([]common.Address, error) {
	return _AccountManager.Contract.GetAssociatedAddresses(&_AccountManager.CallOpts, userPublicKey)
}

// GetAssociatedAddresses is a free data retrieval call binding the contract method 0xc7c04cc0.
//
// Solidity: function getAssociatedAddresses(userPublicKey address) constant returns(address[])
func (_AccountManager *AccountManagerCallerSession) GetAssociatedAddresses(userPublicKey common.Address) ([]common.Address, error) {
	return _AccountManager.Contract.GetAssociatedAddresses(&_AccountManager.CallOpts, userPublicKey)
}

// GetUUID is a free data retrieval call binding the contract method 0x8c01069f.
//
// Solidity: function getUUID(userPublicKey address) constant returns(bytes32)
func (_AccountManager *AccountManagerCaller) GetUUID(opts *bind.CallOpts, userPublicKey common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getUUID", userPublicKey)
	return *ret0, err
}

// GetUUID is a free data retrieval call binding the contract method 0x8c01069f.
//
// Solidity: function getUUID(userPublicKey address) constant returns(bytes32)
func (_AccountManager *AccountManagerSession) GetUUID(userPublicKey common.Address) ([32]byte, error) {
	return _AccountManager.Contract.GetUUID(&_AccountManager.CallOpts, userPublicKey)
}

// GetUUID is a free data retrieval call binding the contract method 0x8c01069f.
//
// Solidity: function getUUID(userPublicKey address) constant returns(bytes32)
func (_AccountManager *AccountManagerCallerSession) GetUUID(userPublicKey common.Address) ([32]byte, error) {
	return _AccountManager.Contract.GetUUID(&_AccountManager.CallOpts, userPublicKey)
}

// GetUUIDString is a free data retrieval call binding the contract method 0xd669582c.
//
// Solidity: function getUUIDString(userPublicKey address) constant returns(string)
func (_AccountManager *AccountManagerCaller) GetUUIDString(opts *bind.CallOpts, userPublicKey common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "getUUIDString", userPublicKey)
	return *ret0, err
}

// GetUUIDString is a free data retrieval call binding the contract method 0xd669582c.
//
// Solidity: function getUUIDString(userPublicKey address) constant returns(string)
func (_AccountManager *AccountManagerSession) GetUUIDString(userPublicKey common.Address) (string, error) {
	return _AccountManager.Contract.GetUUIDString(&_AccountManager.CallOpts, userPublicKey)
}

// GetUUIDString is a free data retrieval call binding the contract method 0xd669582c.
//
// Solidity: function getUUIDString(userPublicKey address) constant returns(string)
func (_AccountManager *AccountManagerCallerSession) GetUUIDString(userPublicKey common.Address) (string, error) {
	return _AccountManager.Contract.GetUUIDString(&_AccountManager.CallOpts, userPublicKey)
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

// RegistrationKey is a free data retrieval call binding the contract method 0x3b197828.
//
// Solidity: function registrationKey() constant returns(address)
func (_AccountManager *AccountManagerCaller) RegistrationKey(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "registrationKey")
	return *ret0, err
}

// RegistrationKey is a free data retrieval call binding the contract method 0x3b197828.
//
// Solidity: function registrationKey() constant returns(address)
func (_AccountManager *AccountManagerSession) RegistrationKey() (common.Address, error) {
	return _AccountManager.Contract.RegistrationKey(&_AccountManager.CallOpts)
}

// RegistrationKey is a free data retrieval call binding the contract method 0x3b197828.
//
// Solidity: function registrationKey() constant returns(address)
func (_AccountManager *AccountManagerCallerSession) RegistrationKey() (common.Address, error) {
	return _AccountManager.Contract.RegistrationKey(&_AccountManager.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x75e36616.
//
// Solidity: function verify(hashedMessage bytes32) constant returns(bool)
func (_AccountManager *AccountManagerCaller) Verify(opts *bind.CallOpts, hashedMessage [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "verify", hashedMessage)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x75e36616.
//
// Solidity: function verify(hashedMessage bytes32) constant returns(bool)
func (_AccountManager *AccountManagerSession) Verify(hashedMessage [32]byte) (bool, error) {
	return _AccountManager.Contract.Verify(&_AccountManager.CallOpts, hashedMessage)
}

// Verify is a free data retrieval call binding the contract method 0x75e36616.
//
// Solidity: function verify(hashedMessage bytes32) constant returns(bool)
func (_AccountManager *AccountManagerCallerSession) Verify(hashedMessage [32]byte) (bool, error) {
	return _AccountManager.Contract.Verify(&_AccountManager.CallOpts, hashedMessage)
}

// AddAddresses is a paid mutator transaction binding the contract method 0x36de7c11.
//
// Solidity: function addAddresses(uuid bytes32, addresses address[]) returns()
func (_AccountManager *AccountManagerTransactor) AddAddresses(opts *bind.TransactOpts, uuid [32]byte, addresses []common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "addAddresses", uuid, addresses)
}

// AddAddresses is a paid mutator transaction binding the contract method 0x36de7c11.
//
// Solidity: function addAddresses(uuid bytes32, addresses address[]) returns()
func (_AccountManager *AccountManagerSession) AddAddresses(uuid [32]byte, addresses []common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AddAddresses(&_AccountManager.TransactOpts, uuid, addresses)
}

// AddAddresses is a paid mutator transaction binding the contract method 0x36de7c11.
//
// Solidity: function addAddresses(uuid bytes32, addresses address[]) returns()
func (_AccountManager *AccountManagerTransactorSession) AddAddresses(uuid [32]byte, addresses []common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.AddAddresses(&_AccountManager.TransactOpts, uuid, addresses)
}

// Recover is a paid mutator transaction binding the contract method 0x7f52ebb2.
//
// Solidity: function recover(oldRPHash bytes32, oldSignedRPHash bytes32, newRPHash bytes32, newSignedRPHash bytes32) returns(bool)
func (_AccountManager *AccountManagerTransactor) Recover(opts *bind.TransactOpts, oldRPHash [32]byte, oldSignedRPHash [32]byte, newRPHash [32]byte, newSignedRPHash [32]byte) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "recover", oldRPHash, oldSignedRPHash, newRPHash, newSignedRPHash)
}

// Recover is a paid mutator transaction binding the contract method 0x7f52ebb2.
//
// Solidity: function recover(oldRPHash bytes32, oldSignedRPHash bytes32, newRPHash bytes32, newSignedRPHash bytes32) returns(bool)
func (_AccountManager *AccountManagerSession) Recover(oldRPHash [32]byte, oldSignedRPHash [32]byte, newRPHash [32]byte, newSignedRPHash [32]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Recover(&_AccountManager.TransactOpts, oldRPHash, oldSignedRPHash, newRPHash, newSignedRPHash)
}

// Recover is a paid mutator transaction binding the contract method 0x7f52ebb2.
//
// Solidity: function recover(oldRPHash bytes32, oldSignedRPHash bytes32, newRPHash bytes32, newSignedRPHash bytes32) returns(bool)
func (_AccountManager *AccountManagerTransactorSession) Recover(oldRPHash [32]byte, oldSignedRPHash [32]byte, newRPHash [32]byte, newSignedRPHash [32]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Recover(&_AccountManager.TransactOpts, oldRPHash, oldSignedRPHash, newRPHash, newSignedRPHash)
}

// Register is a paid mutator transaction binding the contract method 0x4da274fd.
//
// Solidity: function register(uuid bytes32, signedPublicKey bytes32, hashedRecoveryPhrase bytes32) returns(bool)
func (_AccountManager *AccountManagerTransactor) Register(opts *bind.TransactOpts, uuid [32]byte, signedPublicKey [32]byte, hashedRecoveryPhrase [32]byte) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "register", uuid, signedPublicKey, hashedRecoveryPhrase)
}

// Register is a paid mutator transaction binding the contract method 0x4da274fd.
//
// Solidity: function register(uuid bytes32, signedPublicKey bytes32, hashedRecoveryPhrase bytes32) returns(bool)
func (_AccountManager *AccountManagerSession) Register(uuid [32]byte, signedPublicKey [32]byte, hashedRecoveryPhrase [32]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Register(&_AccountManager.TransactOpts, uuid, signedPublicKey, hashedRecoveryPhrase)
}

// Register is a paid mutator transaction binding the contract method 0x4da274fd.
//
// Solidity: function register(uuid bytes32, signedPublicKey bytes32, hashedRecoveryPhrase bytes32) returns(bool)
func (_AccountManager *AccountManagerTransactorSession) Register(uuid [32]byte, signedPublicKey [32]byte, hashedRecoveryPhrase [32]byte) (*types.Transaction, error) {
	return _AccountManager.Contract.Register(&_AccountManager.TransactOpts, uuid, signedPublicKey, hashedRecoveryPhrase)
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
	Uuid [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAccountRecovered is a free log retrieval operation binding the contract event 0xd67568f31e693c2a76b51fa11a1edf173364e7f7c90f3bf34f802b4a49b8344d.
//
// Solidity: event AccountRecovered(uuid bytes32)
func (_AccountManager *AccountManagerFilterer) FilterAccountRecovered(opts *bind.FilterOpts) (*AccountManagerAccountRecoveredIterator, error) {

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return &AccountManagerAccountRecoveredIterator{contract: _AccountManager.contract, event: "AccountRecovered", logs: logs, sub: sub}, nil
}

// WatchAccountRecovered is a free log subscription operation binding the contract event 0xd67568f31e693c2a76b51fa11a1edf173364e7f7c90f3bf34f802b4a49b8344d.
//
// Solidity: event AccountRecovered(uuid bytes32)
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

