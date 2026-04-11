package executor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
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

func (e *Executor) TransferERC20(ctx context.Context, token common.Address, to common.Address, amount *big.Int) error {
	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("get chain id: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.solverKey, chainID)
	if err != nil {
		return fmt.Errorf("create transactor: %w", err)
	}

	erc20, err := bindings.NewIERC20(token, e.client)
	if err != nil {
		return fmt.Errorf("bind erc20 at %s: %w", token.Hex(), err)
	}

	balance, err := erc20.BalanceOf(&bind.CallOpts{Context: ctx}, e.solverAddr)
	if err != nil {
		return fmt.Errorf("check balance: %w", err)
	}
	if balance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient solver balance: have %s, need %s (token %s)", balance, amount, token.Hex())
	}

	tx, err := erc20.Transfer(auth, to, amount)
	if err != nil {
		return fmt.Errorf("transfer %s to %s: %w", amount, to.Hex(), err)
	}

	fmt.Printf("ERC20 transfer tx sent: %s (token=%s, to=%s, amount=%s)\n", tx.Hash().Hex(), token.Hex(), to.Hex(), amount)
	return nil
}

func (e *Executor) WrapETH(ctx context.Context, wethAddr common.Address, amount *big.Int) error {
	chainID, err := e.client.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("get chain id: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(e.solverKey, chainID)
	if err != nil {
		return fmt.Errorf("create transactor: %w", err)
	}

	ethBalance, err := e.client.BalanceAt(ctx, e.solverAddr, nil)
	if err != nil {
		return fmt.Errorf("check ETH balance: %w", err)
	}
	if ethBalance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient ETH for wrapping: have %s, need %s", ethBalance, amount)
	}

	weth, err := bindings.NewIWETH(wethAddr, e.client)
	if err != nil {
		return fmt.Errorf("bind WETH at %s: %w", wethAddr.Hex(), err)
	}

	tx, err := weth.Deposit(auth, amount)
	if err != nil {
		return fmt.Errorf("WETH deposit: %w", err)
	}

	fmt.Printf("WETH wrap tx sent: %s (weth=%s, amount=%s)\n", tx.Hash().Hex(), wethAddr.Hex(), amount)
	return nil
}

func (e *Executor) SolverAddress() common.Address {
	return e.solverAddr
}
