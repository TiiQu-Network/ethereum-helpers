package key

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	privateKeyString = "a8ecbc89f16397032fa3ed158067af048db0b9b2bdc3791b47c219731b8f0284"
	publicKeyString  = "045358b81620342f4141e60acd4466835a90b28e2b895855aba7b664d9dc2c539d9f328ece14f6633ff0467f1c28ac6f357ba4c5a6beded5043346aa0946990dfc"
	addressString    = "4e3e41c5ff3448c4646f68dc363fa0d85cafcaff"
	message          = "Hello this is a message"
	signature        = "11d5ec47dde2893ca0369ff24fe285363c09710cd0e354801ac0b58472c754a4128532f13fb625ea2e8386e7e1499e36f0c27aac37079ad819e030b9c4bdd466"
)

var (
	privateKey = FromString(privateKeyString)
)

func TestPrivateKey_String(t *testing.T) {
	if privateKeyString != privateKey.String {
		t.Error("Private Key String doesn't match")
	}
}

func TestPrivateKey_PublicKey(t *testing.T) {
	if publicKeyString != privateKey.PublicKey() {
		t.Error("Public Key String doesn't match")
	}
}

func TestPrivateKey_Address(t *testing.T) {
	if addressString != privateKey.Address() {
		t.Error("Address String doesn't match")
	}
}

func TestPrivateKey_Sign(t *testing.T) {
	if signature != privateKey.Sign(message) {
		t.Error("Signature doesn't match")
	}
}

func TestPrivateKey(t *testing.T) {
	t.Skip()
	t.Log("*** Setup Private Key ***")
	privateKey := FromString(privateKeyString)

	t.Log("Private Key :", privateKey.String)
	t.Log("Public Key :", privateKey.PublicKey())
	t.Log("Public Key Bytes :", privateKey.EcdsaPubBytes)
	t.Log("Address :", privateKey.Address())
	t.Log("------------------------")

	t.Log("*** Handle Message ***")
	t.Log("Message :", message)
	signature := privateKey.Sign(message)
	t.Log("Signature of message :", signature)
	t.Log("------------------------")

	t.Log("*** Verify Signature ***")

	decode, err := hex.DecodeString(signature)
	t.Log("String decode error :", err)
	t.Log("Decoded signature string :", decode)
	hash := crypto.Keccak256([]byte(message))
	t.Log("Message hash :", hash)
	result := crypto.VerifySignature(privateKey.EcdsaPubBytes, hash, decode)
	t.Log("Verify signature result :", result)
}
