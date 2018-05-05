package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/TiiQu-Network/ethereum-helpers/contracts"
)

func (c *Client) GetUserUUID(contract string, address common.Address, addressSignature []byte) ([32]byte, error) {
	var err error

	conn := c.EtherClient

	accountManager, err := contracts.NewAccountManager(common.HexToAddress(contract), conn)
	if err != nil {
		return *new([32]byte), err
	}

	var r, s [32]byte
	var v uint8

	copy(r[:], addressSignature[0:32])
	copy(s[:], addressSignature[32:64])
	v = addressSignature[64]

	// [R || S || V] format where V is 0 or 1.
	return accountManager.GetUUID(nil, address, r, s, v)
}
