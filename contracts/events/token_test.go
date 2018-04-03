package events

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestLogTransferToken(t *testing.T) {
	var data = `{
    	"logIndex": "0x0",
      	"transactionIndex": "0x0",
      	"transactionHash": "0xd79f2b7b08dab487b51b9f7557ce736b069195677e445926d58b9b174614d051",
      	"blockHash": "0x0248c44667ce668023a7e17b60ea2c67f2911b33630004e677fe259bea7e20a7",
      	"blockNumber": "0x21",
      	"address": "0xef03118e3e60d9003b6c622a2e94c39ebb6985f2",
      	"data": "0x00000000000000000000000000000000000000000000000000000000000f4240",
      	"topics":[
			"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
         	"0x0000000000000000000000000f21f6fb13310ac0e17205840a91da93119efbec",
         	"0x000000000000000000000000d435bbbaa004889f95f54e8232575d87793b42df"
		],
		"type": "mined"
	}`

	log, parseErr := ParseLog(data)
	if parseErr != nil {
		t.Error("Failed to parse JSON data into Ethereum Log :", parseErr)
	}

	var expected = EventTokenTransfer{
		From:  common.HexToAddress("0x0f21F6fB13310AC0E17205840a91dA93119efbec"),
		To:    common.HexToAddress("0xd435bbBAA004889F95F54E8232575d87793B42df"),
		Value: big.NewInt(1000000),
	}

	var o EventTokenTransfer
	err := o.Unmarshal(log)
	if err != nil {
		t.Error("Failed to unmarshal Ethereum log :", err)
	}

	if !reflect.DeepEqual(o, expected) {
		t.Error("Unmarshalled Ethereum log does not match expected Event struct")
	}
}
