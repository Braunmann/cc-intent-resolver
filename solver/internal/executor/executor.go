package executor

import (
	"context"
	"crypto/ecdsa"
	bindings "solver/internal/bindings"
	"solver/internal/store"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Executor struct {
	client     *ethclient.Client
	intentHub  *bindings.IntentHub
	solverKey  *ecdsa.PrivateKey
	solverAddr common.Address
	store      *store.IntentStore
}

func NewExecutor(rpcURL string, contractAddr common.Address, solverKey string, store *store.IntentStore) (*Executor, error) {
	solverKey = strings.TrimPrefix(solverKey, "0x")
	solverKeyECDSA, err := crypto.HexToECDSA(solverKey)
	if err != nil {
		return nil, err
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	intentHub, err := bindings.NewIntentHub(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &Executor{
		client:     client,
		intentHub:  intentHub,
		solverKey:  solverKeyECDSA,
		solverAddr: crypto.PubkeyToAddress(solverKeyECDSA.PublicKey),
		store:      store,
	}, nil
}

func (e *Executor) ExecuteFulfill(ctx context.Context, intent store.Intent) error {
	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.solverKey, chainID)
	if err != nil {
		return err
	}

	_, err = e.intentHub.FulfillIntent(auth, intent.ID, e.solverAddr)
	if err != nil {
		return err
	}

	return nil
}

func (e *Executor) ExecuteSettle(ctx context.Context, intent store.Intent) error {
	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.solverKey, chainID)
	if err != nil {
		return err
	}

	_, err = e.intentHub.SettleIntent(auth, intent.ID)
	if err != nil {
		return err
	}

	return nil
}
