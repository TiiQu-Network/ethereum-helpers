package test

import (
	"testing"
	"github.com/TiiQu-Network/ethereum-helpers/key"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"os"
	"github.com/TiiQu-Network/ethereum-helpers/client"
	"github.com/satori/go.uuid"
	"github.com/TiiQu-Network/ethereum-helpers/contracts"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestAccountManager(t *testing.T) {

	userKey := key.New()
	opts := bind.NewKeyedTransactor(userKey.ECDSA())

	rpc, defined := os.LookupEnv("RPC_ADDRESS")
	if !defined {
		rpc = "http://127.0.0.1:7545"
	}

	ethClient := client.Dial(rpc)

	addr, _, accountManagerContract, err := contracts.DeployAccountManager(opts, ethClient.EtherClient, userKey.AddressEth())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Contract deployed to ", addr.String())


	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	userTransactOpts := bind.NewKeyedTransactor(privateKey)

	sig := userKey.SignToBytes(userKey.AddressBytes())

	id := uuid.Must(uuid.NewV4())
	t.Log("uuid", id.String())

	var uuidBytes, recovery [32]byte
	copy(uuidBytes[:], id.Bytes())

	recovery = *new([32]byte)

	transaction, err := accountManagerContract.Register(userTransactOpts, uuidBytes, recovery, sig.R, sig.S, sig.V)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("User registered ", transaction.String())

}
