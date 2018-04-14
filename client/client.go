package client

import (
	"context"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/Samyoul/go-handle"
	"github.com/TiiQu-Network/ethereum-helpers/contracts"
)

type Client struct {
	errorHandler *handlers.ErrorHandler
	EtherClient  *ethclient.Client
	RpcClient    *rpc.Client
}

func Dial(rawUrl string) (*Client) {

	handler := new(handlers.ErrorHandler)

	// Make connection to Ethereum client
	ethClient, err := ethclient.Dial(rawUrl)
	handler.Handle(err, "Failed to connect to the Ethereum network:")

	// Make connection to the RPC client
	newRpcClient, err := rpc.Dial(rawUrl)
	handler.Handle(err, "Failed to connect to the Ethereum network:")

	return &Client{
		errorHandler: handler,
		EtherClient:  ethClient,
		RpcClient:    newRpcClient,
	}
}

// Internal call function
func (c Client) rpcCall(method string, args ...interface{}) (string) {
	var output string

	err := c.RpcClient.Call(&output, method, args...)
	c.errorHandler.Handle(err, "RPC call failed - Method: %s", method)

	return output
}

/**
 *  RPC Requests
 */

// Get client version
func (c Client) ClientVersion() (string) {
	return c.rpcCall("web3_clientVersion")
}

// Get net version
func (c Client) NetVersion() (string) {
	return c.rpcCall("net_version")
}

// Get latest block number
func (c Client) BlockNumber() (*big.Int) {
	response := c.rpcCall("eth_blockNumber")

	// Convert the hex response data into an int
	blockNumberInt, err := strconv.ParseInt(response, 0, 64)
	c.errorHandler.Handle(err, "Failed to parse blockNumber from hex to int:")

	// Convert int into big.Int for use with other method calls
	return big.NewInt(blockNumberInt)
}

// Get transaction count
func (c Client) TransactionCount(hexAddress string) (*big.Int) {
	response := c.rpcCall("eth_getTransactionCount", hexAddress)

	// Convert the hex response data into an int
	count, err := strconv.ParseInt(response, 0, 64)
	c.errorHandler.Handle(err, "Failed to parse blockNumber from hex to int:")

	// Convert int into big.Int for use with other method calls
	return big.NewInt(count)
}

/**
 *  Ether Requests
 */

func (c Client) BalanceAt(hexAddress string) (*big.Int) {
	blockNumber := c.BlockNumber()

	bal, err := c.EtherClient.BalanceAt(context.TODO(), common.HexToAddress(hexAddress), blockNumber)
	c.errorHandler.Handle(err, "Failed to get balance:")

	return bal
}

func (c *Client) GetContractBalance(contract string, wallet string) (*big.Int, error) {
	var err error

	conn := c.EtherClient

	token, err := contracts.NewTokenCaller(common.HexToAddress(contract), conn)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(wallet)
	if err != nil {
		return nil, err
	}

	balance, err := token.BalanceOf(nil, address)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
