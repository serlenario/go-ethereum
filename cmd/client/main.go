package main

import (
	"github/serlenario/go-gRPC-ethereum/internal/client"
	"github/serlenario/go-gRPC-ethereum/internal/walletgen"
	"log"
)

func main() {
	fileName := "addresses.txt"

	err := walletgen.GenerateWallets(fileName, 1000)
	if err != nil {
		log.Printf("Failed to generate wallets: %v", err)
		return
	}

	client.RunGetAccount()

	client.RunGetAccounts()
}
