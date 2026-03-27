package chain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"solver/internal/store"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChainListener struct {
	rpcURL          string
	contractAddress common.Address
	store           *store.IntentStore
	client          *ethclient.Client
	contractABI     *abi.ABI
}

func LoadABI(path string) (*abi.ABI, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var artifact struct {
		ABI json.RawMessage `json:"abi"`
	}

	if err := json.Unmarshal(data, &artifact); err != nil {
		return nil, err
	}

	parsed, err := abi.JSON(bytes.NewReader(artifact.ABI))
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

func NewChainListener(rpcURL string, contractAddr common.Address, store *store.IntentStore) (*ChainListener, error) {
	contractABI, err := LoadABI("internal/chain/IntentHub.json")
	if err != nil {
		return nil, err
	}

	ethclient, err := ethclient.DialContext(context.Background(), rpcURL)
	if err != nil {
		return nil, err
	}

	return &ChainListener{
		rpcURL:          rpcURL,
		contractAddress: contractAddr,
		store:           store,
		client:          ethclient,
		contractABI:     contractABI,
	}, nil
}

func (l *ChainListener) Listen(ctx context.Context) error {
	logsChan := make(chan types.Log)

	filterQuery := ethereum.FilterQuery{
		Addresses: []common.Address{l.contractAddress},
		Topics: [][]common.Hash{
			{
				l.contractABI.Events["IntentCreated"].ID,
				l.contractABI.Events["IntentFulfilled"].ID,
				l.contractABI.Events["IntentSettled"].ID,
				l.contractABI.Events["IntentCancelled"].ID,
			},
		},
	}

	sub, err := l.client.SubscribeFilterLogs(ctx, filterQuery, logsChan)
	if err != nil {
		return err
	}

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			return err
		case <-ctx.Done():
			return nil
		case log := <-logsChan:
			if err := l.handleLog(log); err != nil {
				return err
			}
		}
	}
}

func (l *ChainListener) handleLog(log types.Log) error {
	switch log.Topics[0] {
	case l.contractABI.Events["IntentCreated"].ID:
		fmt.Println("IntentCreated")
	case l.contractABI.Events["IntentFulfilled"].ID:
		fmt.Println("IntentFulfilled")
	case l.contractABI.Events["IntentSettled"].ID:
		fmt.Println("IntentSettled")
	case l.contractABI.Events["IntentCancelled"].ID:
		fmt.Println("IntentCancelled")
	}
	fmt.Println(log)
	return nil
}
