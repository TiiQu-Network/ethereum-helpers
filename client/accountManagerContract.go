package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/TiiQu-Network/ethereum-helpers/contracts"
)

func (c *Client) GetUserUUID(contract string, address common.Address) (string, error) {
	var err error

	conn := c.EtherClient

	accountManager, err := contracts.NewAccountManager(common.HexToAddress(contract), conn)
	if err != nil {
		return "", err
	}

	return accountManager.GetUUIDString(nil, address)
}
