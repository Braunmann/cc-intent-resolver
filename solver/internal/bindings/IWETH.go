package bindings

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const iwethABI = `[{"constant":false,"inputs":[],"name":"deposit","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"wad","type":"uint256"}],"name":"withdraw","outputs":[],"type":"function"}]`

type IWETH struct {
	contract *bind.BoundContract
}

func NewIWETH(address common.Address, client *ethclient.Client) (*IWETH, error) {
	parsed, err := abi.JSON(strings.NewReader(iwethABI))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(address, parsed, client, client, client)
	return &IWETH{contract: contract}, nil
}

func (w *IWETH) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	opts.Value = amount
	return w.contract.Transact(opts, "deposit")
}
