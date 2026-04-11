// Code generated - DO NOT EDIT.
// This is a minimal ERC20 binding for transfer functionality.

package bindings

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

const ierc20ABI = `[{"constant":false,"inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"type":"function"},{"constant":true,"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"type":"function"}]`

type IERC20 struct {
	contract *bind.BoundContract
}

func NewIERC20(address common.Address, client *ethclient.Client) (*IERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ierc20ABI))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(address, parsed, client, client, client)
	return &IERC20{contract: contract}, nil
}

func (e *IERC20) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return e.contract.Transact(opts, "transfer", to, amount)
}

func (e *IERC20) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := e.contract.Call(opts, &out, "balanceOf", account)
	if err != nil {
		return nil, err
	}
	return out[0].(*big.Int), nil
}
