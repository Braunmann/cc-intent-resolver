package config

import (
	"encoding/json"
	"os"
)

type ChainConfig struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	RegistryKey string `json:"registryKey"`
}

type AppConfig struct {
	Chains        []ChainConfig                `json:"chains"`
	Slippage      float64                      `json:"slippage"`
	TokenDecimals map[string]int               `json:"tokenDecimals"`
	TokenPriceKey map[string]string            `json:"tokenPriceKey"`
	Tokens        map[string]map[string]string `json:"tokens"`
}

var App = AppConfig{
	Chains: []ChainConfig{
		{ID: 11155111, Name: "Sepolia", RegistryKey: "sepolia"},
		{ID: 11155420, Name: "OP Sepolia", RegistryKey: "opSepolia"},
	},
	Slippage: 0.05,
	TokenDecimals: map[string]int{
		"WETH": 18,
		"DAI":  18,
		"USDC": 6,
	},
	TokenPriceKey: map[string]string{
		"WETH": "weth",
		"DAI":  "dai",
		"USDC": "usdc",
	},
	Tokens: map[string]map[string]string{
		"sepolia": {
			"WETH": "0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9",
			"USDC": "0x1c7D4B196Cb0C6fb792C6b106d5d4d3A34566e8b",
			"DAI":  "0x3e622317f8C93f7328350cF0B56d9eD4C620C5d6",
		},
		"opSepolia": {
			"WETH": "0x4200000000000000000000000000000000000006",
			"USDC": "0x5fd84259d66Cd46628340c017a0563B6c4e4414c",
			"DAI":  "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1",
		},
	},
}

func init() {
	path := os.Getenv("TOKEN_CONFIG_PATH")
	if path == "" {
		return
	}
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	var override map[string]map[string]string
	if err := json.NewDecoder(f).Decode(&override); err != nil {
		return
	}
	for key, tokenMap := range override {
		if len(tokenMap) > 0 {
			App.Tokens[key] = tokenMap
		}
	}
}
