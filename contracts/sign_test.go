package contracts

import (
	"context"
	"os"
	"testing"

	"github.com/TiiQu-Network/ethereum-helpers/client"
	"github.com/TiiQu-Network/ethereum-helpers/key"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TestSign(t *testing.T) {
	userKey := key.New()
	opts := bind.NewKeyedTransactor(userKey.ECDSA())
	ethClient := client.Dial(os.Getenv("RPC_ADDRESS"))
	userCallOpts := &bind.CallOpts{
		Pending: false,
		From:    userKey.AddressEth(),
		Context: context.TODO(),
	}

	addr, _, signContract, err := DeploySign(opts, ethClient.EtherClient, userKey.AddressEth())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addr)

	sig := userKey.SignToBytes(userKey.AddressBytes())

	result, err := signContract.Verify(userCallOpts, sig.R, sig.S, sig.V)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
