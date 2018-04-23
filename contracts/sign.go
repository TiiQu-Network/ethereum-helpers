package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"strings"
)

// SignABI is the input ABI used to generate the binding from.
const SignABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getVettingAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"VettingAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SignBin is the compiled bytecode used for deploying new contracts.
const SignBin = `0x6060604052341561000f57600080fd5b6040516020806102208339810160405280805160008054600160a060020a03909216600160a060020a031990921691909117905550506101cc806100546000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631288d75481146100505780633dbdd9051461007f575b600080fd5b341561005b57600080fd5b6100636100b2565b604051600160a060020a03909116815260200160405180910390f35b341561008a57600080fd5b61009e60043560243560ff604435166100c1565b604051901515815260200160405180910390f35b600054600160a060020a031690565b600061010a33604051600160a060020a03919091166c01000000000000000000000000028152601401604051908190039020600054869086908690600160a060020a0316610112565b949350505050565b600081600160a060020a03166001878588886040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f1151561018357600080fd5b505060206040510351600160a060020a03161496955050505050505600a165627a7a7230582090d515945526a290932ba8b90c3ee2c8a081ab800e06cb07b9d5d98a235f19180029`

// DeploySign deploys a new Ethereum contract, binding an instance of Sign to it.
func DeploySign(auth *bind.TransactOpts, backend bind.ContractBackend, VettingAddress common.Address) (common.Address, *types.Transaction, *Sign, error) {
	parsed, err := abi.JSON(strings.NewReader(SignABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SignBin), backend, VettingAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sign{SignCaller: SignCaller{contract: contract}, SignTransactor: SignTransactor{contract: contract}}, nil
}

// Sign is an auto generated Go binding around an Ethereum contract.
type Sign struct {
	SignCaller     // Read-only binding to the contract
	SignTransactor // Write-only binding to the contract
}

// SignCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignSession struct {
	Contract     *Sign             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignCallerSession struct {
	Contract *SignCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignTransactorSession struct {
	Contract     *SignTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignRaw struct {
	Contract *Sign // Generic contract binding to access the raw methods on
}

// SignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignCallerRaw struct {
	Contract *SignCaller // Generic read-only contract binding to access the raw methods on
}

// SignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignTransactorRaw struct {
	Contract *SignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSign creates a new instance of Sign, bound to a specific deployed contract.
func NewSign(address common.Address, backend bind.ContractBackend) (*Sign, error) {
	contract, err := bindSign(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sign{SignCaller: SignCaller{contract: contract}, SignTransactor: SignTransactor{contract: contract}}, nil
}

// NewSignCaller creates a new read-only instance of Sign, bound to a specific deployed contract.
func NewSignCaller(address common.Address, caller bind.ContractCaller) (*SignCaller, error) {
	contract, err := bindSign(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SignCaller{contract: contract}, nil
}

// NewSignTransactor creates a new write-only instance of Sign, bound to a specific deployed contract.
func NewSignTransactor(address common.Address, transactor bind.ContractTransactor) (*SignTransactor, error) {
	contract, err := bindSign(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SignTransactor{contract: contract}, nil
}

// bindSign binds a generic wrapper to an already deployed contract.
func bindSign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (Sign *SignRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return Sign.Contract.SignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (Sign *SignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Sign.Contract.SignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (Sign *SignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return Sign.Contract.SignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (Sign *SignCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return Sign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (Sign *SignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Sign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (Sign *SignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return Sign.Contract.contract.Transact(opts, method, params...)
}

// GetVettingAddress is a free data retrieval call binding the contract method 0x1288d754.
//
// Solidity: function getVettingAddress() constant returns(address)
func (Sign *SignCaller) GetVettingAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := Sign.contract.Call(opts, out, "getVettingAddress")
	return *ret0, err
}

// GetVettingAddress is a free data retrieval call binding the contract method 0x1288d754.
//
// Solidity: function getVettingAddress() constant returns(address)
func (Sign *SignSession) GetVettingAddress() (common.Address, error) {
	return Sign.Contract.GetVettingAddress(&Sign.CallOpts)
}

// GetVettingAddress is a free data retrieval call binding the contract method 0x1288d754.
//
// Solidity: function getVettingAddress() constant returns(address)
func (Sign *SignCallerSession) GetVettingAddress() (common.Address, error) {
	return Sign.Contract.GetVettingAddress(&Sign.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x3dbdd905.
//
// Solidity: function verify(r bytes32, s bytes32, v uint8) constant returns(bool)
func (Sign *SignCaller) Verify(opts *bind.CallOpts, r [32]byte, s [32]byte, v uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := Sign.contract.Call(opts, out, "verify", r, s, v)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x3dbdd905.
//
// Solidity: function verify(r bytes32, s bytes32, v uint8) constant returns(bool)
func (Sign *SignSession) Verify(r [32]byte, s [32]byte, v uint8) (bool, error) {
	return Sign.Contract.Verify(&Sign.CallOpts, r, s, v)
}

// Verify is a free data retrieval call binding the contract method 0x3dbdd905.
//
// Solidity: function verify(r bytes32, s bytes32, v uint8) constant returns(bool)
func (Sign *SignCallerSession) Verify(r [32]byte, s [32]byte, v uint8) (bool, error) {
	return Sign.Contract.Verify(&Sign.CallOpts, r, s, v)
}

// SignatureABI is the input ABI used to generate the binding from.
const SignatureABI = "[]"

// SignatureBin is the compiled bytecode used for deploying new contracts.
const SignatureBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a7230582005b64630f283c9b3b17d450fb5094990ad243e8f370b7e81d429447a0ffb88c90029`

// DeploySignature deploys a new Ethereum contract, binding an instance of Signature to it.
func DeploySignature(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Signature, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SignatureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Signature{SignatureCaller: SignatureCaller{contract: contract}, SignatureTransactor: SignatureTransactor{contract: contract}}, nil
}

// Signature is an auto generated Go binding around an Ethereum contract.
type Signature struct {
	SignatureCaller     // Read-only binding to the contract
	SignatureTransactor // Write-only binding to the contract
}

// SignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureSession struct {
	Contract     *Signature        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureCallerSession struct {
	Contract *SignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureTransactorSession struct {
	Contract     *SignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureRaw struct {
	Contract *Signature // Generic contract binding to access the raw methods on
}

// SignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureCallerRaw struct {
	Contract *SignatureCaller // Generic read-only contract binding to access the raw methods on
}

// SignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureTransactorRaw struct {
	Contract *SignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignature creates a new instance of Signature, bound to a specific deployed contract.
func NewSignature(address common.Address, backend bind.ContractBackend) (*Signature, error) {
	contract, err := bindSignature(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Signature{SignatureCaller: SignatureCaller{contract: contract}, SignatureTransactor: SignatureTransactor{contract: contract}}, nil
}

// NewSignatureCaller creates a new read-only instance of Signature, bound to a specific deployed contract.
func NewSignatureCaller(address common.Address, caller bind.ContractCaller) (*SignatureCaller, error) {
	contract, err := bindSignature(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureCaller{contract: contract}, nil
}

// NewSignatureTransactor creates a new write-only instance of Signature, bound to a specific deployed contract.
func NewSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*SignatureTransactor, error) {
	contract, err := bindSignature(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SignatureTransactor{contract: contract}, nil
}

// bindSignature binds a generic wrapper to an already deployed contract.
func bindSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (Signature *SignatureRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return Signature.Contract.SignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (Signature *SignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Signature.Contract.SignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (Signature *SignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return Signature.Contract.SignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (Signature *SignatureCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return Signature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (Signature *SignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return Signature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (Signature *SignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return Signature.Contract.contract.Transact(opts, method, params...)
}
