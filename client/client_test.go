package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/TiiQu-Network/ethereum-helpers/contracts/events"
	"github.com/TiiQu-Network/ethereum-helpers/key"
	"github.com/Samyoul/go-handle"
	"math/big"
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

	account := "0x4841F78f190B7436a0048332d871a5cF7C8A8636"

	// Get balance of account
	bal := ethClient.BalanceAt(account)
	fmt.Println("Balance is : " + bal.String())

	// Get transaction count
	txCount := ethClient.TransactionCount(account)
	fmt.Println("Transaction count is : " + txCount.String())

	privateKey := key.New()
	bind.NewKeyedTransactor(privateKey.ECDSA())

	fmt.Println("Signature: " + events.TransferTokenSignature.String())

	balance, err := ethClient.GetContractBalance("0x2c3c1f05187dba7a5f2dd47dca57281c4d4f183f", account)
	handler.Handle(err, "Balance Error")
	fmt.Println("balance is " + balance.String())

	logs, err := ethClient.GetContractLogs("0x2c3c1f05187dba7a5f2dd47dca57281c4d4f183f", big.NewInt(5300000), big.NewInt(5400000))
	handler.Handle(err, "Logs Error")
	printLogs(logs)
}


func printLogs(logs []ContractLog) {
	fmt.Println("Found logs :", len(logs))
	fmt.Println("----------------------------------------------------------------")

	for index, element := range logs {

		fmt.Println("Event :", index)
		fmt.Println("- TX Hash :", element.TxHash)
		fmt.Println("- Block :", element.Block)
		fmt.Println("- From :", element.From)
		fmt.Println("- To :", element.To)
		fmt.Println("- Removed :", element.Removed)
		fmt.Println("- Value :", element.Value.String())

		fmt.Println("----------------------------------------------------------------")
	}

}
