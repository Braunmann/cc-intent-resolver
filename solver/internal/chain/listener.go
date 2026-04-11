package chain

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"solver/internal/config"
	"solver/internal/executor"
	"solver/internal/store"
	"strings"

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
	sourceExecutor  *executor.Executor
	targetExecutor  *executor.Executor
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

func NewChainListener(rpcURL string, contractAddr common.Address, store *store.IntentStore, sourceExecutor *executor.Executor, targetExecutor *executor.Executor) (*ChainListener, error) {
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
		sourceExecutor:  sourceExecutor,
		targetExecutor:  targetExecutor,
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

func (l *ChainListener) handleIntentCreated(log types.Log) error {
	event := make(map[string]interface{})
	if err := l.contractABI.UnpackIntoMap(event, "IntentCreated", log.Data); err != nil {
		return err
	}

	intentId := log.Topics[1]
	maker := common.BytesToAddress(log.Topics[2].Bytes())

	intent := store.Intent{
		ID:              intentId,
		Maker:           maker,
		InputToken:      event["inputToken"].(common.Address),
		InputAmount:     event["inputAmount"].(*big.Int),
		OutputToken:     event["outputToken"].(common.Address),
		MinOutputAmount: event["minOutputAmount"].(*big.Int),
		TargetChainId:   event["targetChainId"].(uint32),
		Recipient:       event["recipient"].(common.Address),
		Deadline:        event["deadline"].(uint64),
		Nonce:           event["nonce"].(uint64),
		Status:          store.IntentStatusCreated,
	}

	if err := l.store.Save(intent); err != nil {
		return err
	}

	go func() {
		ctx := context.Background()

		// Step 1: If outputToken is WETH, wrap ETH → WETH on target chain
		if isWETH(intent.OutputToken) {
			fmt.Printf("Wrapping ETH → WETH on target chain (token=%s, amount=%s)\n",
				intent.OutputToken.Hex(), intent.MinOutputAmount)
			if err := l.targetExecutor.WrapETH(ctx, intent.OutputToken, intent.MinOutputAmount); err != nil {
				fmt.Println("wrap ETH error:", err)
				return
			}
		}

		// Step 2: Transfer outputToken to recipient on target chain
		if err := l.targetExecutor.TransferERC20(ctx, intent.OutputToken, intent.Recipient, intent.MinOutputAmount); err != nil {
			fmt.Println("cross-chain transfer error:", err)
			return
		}
		fmt.Printf("Cross-chain transfer done: %s -> %s (token=%s, amount=%s)\n",
			l.targetExecutor.SolverAddress().Hex(), intent.Recipient.Hex(), intent.OutputToken.Hex(), intent.MinOutputAmount)

		// Step 3: Fulfill intent on source chain
		if err := l.sourceExecutor.ExecuteFulfill(ctx, intent); err != nil {
			fmt.Println("fulfill error:", err)
		}
	}()

	fmt.Printf("IntentCreated: %v\n", event)
	return nil
}

func (l *ChainListener) handleIntentFulfilled(log types.Log) error {
	intentId := log.Topics[1]
	solverAddr := common.BytesToAddress(log.Topics[2].Bytes())

	intent, ok := l.store.Get(intentId)
	if !ok {
		return fmt.Errorf("intent not found: %s", intentId.Hex())
	}

	intent.Solver = solverAddr
	intent.Status = store.IntentStatusFulfilled

	if err := l.store.Save(intent); err != nil {
		return err
	}

	go func() {
		// Settle intent on source chain (claim locked input tokens)
		if err := l.sourceExecutor.ExecuteSettle(context.Background(), intent); err != nil {
			fmt.Println("settle error:", err)
		}
	}()

	fmt.Println("IntentFulfilled", intentId, intent)
	return nil
}

func (l *ChainListener) handleIntentSettled(log types.Log) error {
	intentId := log.Topics[1]

	intent, ok := l.store.Get(intentId)
	if !ok {
		return fmt.Errorf("intent not found: %s", intentId.Hex())
	}

	intent.Status = store.IntentStatusSettled

	if err := l.store.Save(intent); err != nil {
		return err
	}

	fmt.Println("IntentSettled", intentId, intent)
	return nil
}

func (l *ChainListener) handleIntentCancelled(log types.Log) error {
	intentId := log.Topics[1]

	intent, ok := l.store.Get(intentId)
	if !ok {
		return fmt.Errorf("intent not found: %s", intentId.Hex())
	}

	intent.Status = store.IntentStatusCancelled

	if err := l.store.Save(intent); err != nil {
		return err
	}

	fmt.Println("IntentCancelled", intentId, intent)
	return nil
}

func isWETH(addr common.Address) bool {
	for _, tokenMap := range config.App.Tokens {
		if wethAddr, ok := tokenMap["WETH"]; ok {
			if strings.EqualFold(addr.Hex(), wethAddr) {
				return true
			}
		}
	}
	return false
}

func (l *ChainListener) handleLog(log types.Log) error {
	switch log.Topics[0] {
	case l.contractABI.Events["IntentCreated"].ID:
		return l.handleIntentCreated(log)
	case l.contractABI.Events["IntentFulfilled"].ID:
		return l.handleIntentFulfilled(log)
	case l.contractABI.Events["IntentSettled"].ID:
		return l.handleIntentSettled(log)
	case l.contractABI.Events["IntentCancelled"].ID:
		return l.handleIntentCancelled(log)
	}
	fmt.Println(log)
	return nil
}
