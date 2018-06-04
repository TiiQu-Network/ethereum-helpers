package client

import (
	"math/big"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/TiiQu-Network/ethereum-helpers/contracts/events"
	"context"
	"github.com/TiiQu-Network/ethereum-helpers/contracts"
)

type ContractLog struct {
	TxHash  string
	Removed bool
	From    string
	To      string
	Value   *big.Int
	Block   uint64
}

func (c *Client) GetContractLogs(contract string, from *big.Int, to *big.Int) ([]ContractLog, error) {
	query := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Topics:    [][]common.Hash{{events.TransferTokenSignature}},
		Addresses: []common.Address{common.HexToAddress(contract)},
	}

	results, err := c.EtherClient.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	res := make([]ContractLog, len(results))

	for index, element := range results {

		event := new(events.EventTokenTransfer)
		event.Unmarshal(&element)

		res[index] = ContractLog{
			TxHash:  element.TxHash.String(),
			Removed: element.Removed,
			From:    event.From.String(),
			To:      event.To.String(),
			Value:   event.Value,
			Block: element.BlockNumber,
		}

	}

	return res, nil
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

