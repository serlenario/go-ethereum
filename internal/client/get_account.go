package client

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"
	"github/serlenario/go-gRPC-ethereum/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGetAccount() {
	privateKey, walletAddress := importPrivateKey()

	inputAddress := getAddressFromUser(walletAddress)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing gRPC connection: %v", err)
		}
	}(conn)

	client := proto.NewAccountServiceClient(conn)

	signatureHex := signMessage(privateKey)

	req := &proto.GetAccountRequest{
		EthereumAddress: inputAddress,
		CryptoSignature: signatureHex,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.GetAccount(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetAccount: %v", err)
	}

	fmt.Printf("Gas Token Balance: %s\n", resp.GetGasTokenBalance())
	fmt.Printf("Wallet Nonce: %d\n", resp.GetWalletNonce())
}

func importPrivateKey() (*ecdsa.PrivateKey, string) {
	privateKeyHex := os.Getenv("ETH_PRIVATE_KEY")
	if privateKeyHex == "" {
		log.Fatalf("Environment variable ETH_PRIVATE_KEY is not set")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Unable to recognize private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Invalid public key")
	}

	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return privateKey, walletAddress
}

func getAddressFromUser(defaultAddress string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Ethereum address (or press Enter to use your own): ")
	inputAddress, _ := reader.ReadString('\n')
	inputAddress = strings.TrimSpace(inputAddress)

	if inputAddress == "" {
		inputAddress = defaultAddress
	}
	return inputAddress
}

func signMessage(privateKey *ecdsa.PrivateKey) string {
	message := []byte("Authenticate to DeNet service")
	msgHash := accounts.TextHash(message)

	signature, err := crypto.Sign(msgHash, privateKey)
	if err != nil {
		log.Fatalf("Failed to sign message: %v", err)
	}
	return fmt.Sprintf("%x", signature)
}
