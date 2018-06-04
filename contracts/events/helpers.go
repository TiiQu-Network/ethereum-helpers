package events

import (
	"context"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"fmt"
)

// MustParseABI parses abi string
func MustParseABI(name string, contractABI string) (abi.ABI, error) {
	a, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		err = errors.New("Can't parse contract abi " + name + " - Details : " + err.Error())
	}
	return a, err
}

// MustHaveEvents ensures that the ABI object has listed events
func MustHaveEvents(e abi.ABI, eventNames ...string) error {
	for _, s := range eventNames {
		if _, ok := e.Events[s]; !ok {
			return errors.New("Contract doesn't have required event : " + s)
		}
	}
	return nil
}

// TODO possibly should live in the client class
// SubscribeSimple is a simple utility function to create events subscription
func MakeSimpleSubscription(ctx context.Context, client *ethclient.Client, query ethereum.FilterQuery) (<-chan types.Log, ethereum.Subscription, error) {
	var events = make(chan types.Log, 5) // batch the events, makes processing quicker
	s, err := client.SubscribeFilterLogs(ctx, query, events)
	return events, s, errors.New("Can't create Ethereum Subscription - Details : " + err.Error())
}

// TODO possibly should live in the client class
func WatchSubscription(logs <-chan types.Log, subscription ethereum.Subscription) {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	errChan := subscription.Err()
	for {
		select {
		case err := <-errChan:
			logger.Error("Logs subscription error", zap.Error(err))
			break
		case l := <-logs:
			logData := fmt.Sprintf(`log: %x %x %x %x %d %x %d`, l.Address, l.Topics, l.Data, l.TxHash, l.TxIndex, l.BlockHash, l.Index)
			logger.Info("new log", zap.String("log", logData))
		}
	}
}

// UnmarshalEvent blockchain log into the event structure
// `dest` must be a pointer to initialized structure
func UnmarshalEvent(dest interface{}, data []byte, event abi.Event) error {

	a := abi.ABI{
		Events: map[string]abi.Event{"e": event},
	}

	err := a.Unpack(dest, "e", data)
	if err != nil {
		err = errors.New("The ABI may not match the contract version - Details : " + err.Error())
	}

	return err
}

// ParseLog parses ethereum log JSON
func ParseLog(data string) (*types.Log, error) {
	var log = new(types.Log)
	err := log.UnmarshalJSON([]byte(data))
	return log, err
}
