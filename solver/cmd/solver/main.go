package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"solver/internal/api"
	"solver/internal/chain"
	"solver/internal/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

type Config struct {
	SepoliaRPC   string
	OPSepoliaRPC string
	ContractAddr common.Address
	APIPort      string
}

func loadConfig() Config {
	godotenv.Load()
	return Config{
		SepoliaRPC:   os.Getenv("SEPOLIA_WSS_URL"),
		OPSepoliaRPC: os.Getenv("OP_SEPOLIA_WSS_URL"),
		ContractAddr: common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")),
		APIPort:      os.Getenv("API_PORT"),
	}
}

func mustCreateListener(rpcURL string, contractAddr common.Address, intentStore *store.IntentStore) *chain.ChainListener {
	listener, err := chain.NewChainListener(rpcURL, contractAddr, intentStore)
	if err != nil {
		fmt.Println("Error creating chain listener:", err)
		os.Exit(1)
	}
	fmt.Println("Listening for intents on", rpcURL)
	return listener
}

func mustCreateServer(apiPort string, intentStore *store.IntentStore) *api.Server {
	server, err := api.NewServer(intentStore, apiPort)
	if err != nil {
		fmt.Println("Error creating server:", err)
		os.Exit(1)
	}
	fmt.Println("API server started on port", apiPort)
	return server
}

func waitForShutdown() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	config := loadConfig()

	intentStore := store.NewIntentStore()

	listener := mustCreateListener(config.SepoliaRPC, config.ContractAddr, intentStore)
	opListener := mustCreateListener(config.OPSepoliaRPC, config.ContractAddr, intentStore)

	server := mustCreateServer(":"+config.APIPort, intentStore)

	ctx, cancel := waitForShutdown()
	defer cancel()

	errChan := make(chan error, 2)

	go func() { errChan <- listener.Listen(ctx) }()
	go func() { errChan <- opListener.Listen(ctx) }()

	go server.Start()

	fmt.Println("solver started")

	select {
	case err := <-errChan:
		fmt.Println("listener error:", err)
		cancel()
	case <-ctx.Done():
		fmt.Println("solver stopped")
	}
}
