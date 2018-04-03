package events

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/TiiQu-Network/ethereum-helpers/contracts"
	"github.com/ethereum/go-ethereum/crypto"
)

// globals
var (
	TokenABI, _ = MustParseABI("Token", contracts.TokenABI)
	TransferTokenSignature common.Hash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

const (
	logTokenTransfer = "Transfer"
)

func init() {
	MustHaveEvents(TokenABI, logTokenTransfer)
}

// LogTokenTransfer returns event
func LogTokenTransfer() abi.Event { return TokenABI.Events[logTokenTransfer] }

// EventTokenTransfer represents Transfer event payload
type EventTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// Unmarshal blockchain log into the event structure
func (e *EventTokenTransfer) Unmarshal(log *types.Log) error {
	err := UnmarshalEvent(e, log.Data, LogTokenTransfer())
	if err != nil {
		return err
	}
	if len(log.Topics) < 3 {
		return errors.New("Can't unmarshal EventTokenTransfer event. Not enough elements in Topic. Want at least: 3, got: " + string(len(log.Topics)))
	}
	e.From = common.BytesToAddress(log.Topics[1].Bytes())
	e.To = common.BytesToAddress(log.Topics[2].Bytes())
	return nil
}
