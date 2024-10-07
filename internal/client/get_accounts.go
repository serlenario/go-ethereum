package client

import (
	"bufio"
	"context"
	"fmt"
	"github/serlenario/go-gRPC-ethereum/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func RunGetAccounts() {

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

	addresses, err := readAddressesFromFile("addresses.txt")
	if err != nil {
		log.Fatalf("Error reading addresses: %v", err)
	}

	tokenAddress := getTokenAddressFromUser()

	addressCount := getAddressCountFromUser(len(addresses))

	fmt.Printf("Executing GetAccounts for %d addresses...\n", addressCount)

	//var wg sync.WaitGroup
	var mu sync.Mutex
	index := 1

	//const maxConcurrency = 10
	//semaphore := make(chan struct{}, maxConcurrency)

	//for _, address := range addresses[:addressCount] {
	//	wg.Add(1)
	//	semaphore <- struct{}{}
	//	go func(addr string) {
	//		defer wg.Done()
	//		defer func() { <-semaphore }()
	//		callGetAccounts(client, []string{addr}, tokenAddress, &index, &mu)
	//	}(address)
	//}
	//
	//wg.Wait()

	callGetAccounts(client, addresses[:addressCount], tokenAddress, &index, &mu)
}

func readAddressesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(file)

	var addresses []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		address := strings.TrimSpace(scanner.Text())
		if address != "" {
			addresses = append(addresses, address)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}

func getTokenAddressFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ERC20 token address: ")
	tokenAddress, _ := reader.ReadString('\n')
	return strings.TrimSpace(tokenAddress)
}

func getAddressCountFromUser(maxCount int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter the number of addresses to process (maximum %d): ", maxCount)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		count, err := strconv.Atoi(input)
		if err != nil || count <= 0 || count > maxCount {
			fmt.Printf("Please enter a number between 1 and %d.\n", maxCount)
			continue
		}
		return count
	}
}

func callGetAccounts(client proto.AccountServiceClient, addresses []string, tokenAddress string, index *int, mu *sync.Mutex) {
	stream, err := client.GetAccounts(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	req := &proto.GetAccountsRequest{
		EthereumAddresses: addresses,
		Erc20TokenAddress: tokenAddress,
	}

	if err := stream.Send(req); err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error closing submission: %v", err)
	}

	start := time.Now()

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		mu.Lock()
		fmt.Printf("%d. Address: %s, Balance: %s\n", *index, resp.GetEthereumAddress(), resp.GetErc20Balance())
		*index++
		mu.Unlock()
	}

	elapsed := time.Since(start)
	fmt.Printf("Time spent: %s\n", elapsed)
}
