package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"solver/internal/chain"
	"solver/internal/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	sepoliaRPC := os.Getenv("SEPOLIA_WSS_URL")
	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	intentStore := store.NewIntentStore()

	listener, err := chain.NewChainListener(sepoliaRPC, contractAddr, intentStore)
	if err != nil {
		fmt.Println("Error creating chain listener:", err)
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	fmt.Println("Listening for intents...")

	if err := listener.Listen(ctx); err != nil {
		fmt.Println("listener error:", err)
	}
}
