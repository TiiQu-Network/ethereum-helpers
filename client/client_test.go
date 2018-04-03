package client

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/TiiQu-Network/ethereum-helpers/contracts/events"
	"github.com/TiiQu-Network/ethereum-helpers/key"
	"github.com/Samyoul/go-handle/handlers"
)

var handler = new(handlers.ErrorHandler)

func Test(t *testing.T) {

	// TODO complete these tests

	// Make connection to Ethereum client
	ethClient := Dial(os.Getenv("RPC_ADDRESS"))

	// Get standard connection data
	version := ethClient.ClientVersion()
	net := ethClient.NetVersion()
	blockNumber := ethClient.BlockNumber()

	fmt.Println(version)
	fmt.Println(net)
	fmt.Println(blockNumber)

	// Get balance of account
	bal := ethClient.BalanceAt("0x0000000000000000000000000000000000000000")

	// Output balance
	fmt.Println("Balance is : " + bal.String())

	privateKey := key.New()
	bind.NewKeyedTransactor(privateKey.ECDSA())

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2607801),
		ToBlock:   big.NewInt(2607804),
		Topics:    [][]common.Hash{{events.TransferTokenSignature}},
		Addresses: []common.Address{common.HexToAddress("0xa74476443119A942dE498590Fe1f2454d7D4aC0d")},
	}

	fmt.Println(query)

	results, err := ethClient.EtherClient.FilterLogs(context.TODO(), query)
	handler.Handle(err, "Filter Error")

	fmt.Println("Found logs :", len(results))
	fmt.Println("----------------------------------------------------------------")

	for index, element := range results {

		event := new(events.EventTokenTransfer)
		event.Unmarshal(&element)
		fmt.Println("Event :", index)
		fmt.Println("- From  :", event.From.String())
		fmt.Println("- To    :", event.To.String())
		fmt.Println("- Value :", event.Value)
		fmt.Println("----------------------------------------------------------------")
	}
}
