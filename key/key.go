package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
  
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/samyoul/go-handle/handlers"
)

type PrivateKey struct {
	String        string
	ecdsa         *ecdsa.PrivateKey
	EcdsaPubBytes []byte
	errorHandler  *handlers.ErrorHandler
}

func New() *PrivateKey {
	handler := new(handlers.ErrorHandler)
	ecdsaPrivateKey, err := crypto.GenerateKey()
	handler.Handle(err, "Failed to generate ECDSA private key")

	return generate(handler, ecdsaPrivateKey)
}

func FromString(key string) *PrivateKey {
	handler := new(handlers.ErrorHandler)
	ecdsaPrivateKey, err := crypto.HexToECDSA(key)
	handler.Handle(err, "Failed to convert hex string to ECDSA")

	return generate(handler, ecdsaPrivateKey)
}

func generate(handler *handlers.ErrorHandler, key *ecdsa.PrivateKey) *PrivateKey {

	keyString := hex.EncodeToString(crypto.FromECDSA(key))

	ecdsaPubBytes := elliptic.Marshal(
		secp256k1.S256(),
		key.PublicKey.X,
		key.PublicKey.Y,
	)

	return &PrivateKey{
		String:        keyString,
		ecdsa:         key,
		EcdsaPubBytes: ecdsaPubBytes,
		errorHandler:  handler,
	}
}

func (p PrivateKey) ECDSA() *ecdsa.PrivateKey {
	return p.ecdsa
}

func (p PrivateKey) PublicKey() string {
	return hex.EncodeToString(p.EcdsaPubBytes)
}

func (p PrivateKey) Address() string {
	return hex.EncodeToString(p.AddressBytes())
}

func (p PrivateKey) AddressBytes() []byte {
	return crypto.Keccak256(p.EcdsaPubBytes[1:])[12:]
}

func (p PrivateKey) AddressSha3() string {
	return hex.EncodeToString(crypto.Keccak256(p.EcdsaPubBytes[1:]))
}

func (p PrivateKey) Sign(message string) string {
	hash := crypto.Keccak256([]byte(message))

	signature, err := crypto.Sign(hash, p.ecdsa)
	p.errorHandler.Handle(err, "Signature error")

	return hex.EncodeToString(signature[:64])
}
