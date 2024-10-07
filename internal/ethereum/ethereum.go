package ethereum

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"math"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func InfuraURL() string {
	infuraURL := os.Getenv("INFURA_URL")
	if infuraURL == "" {
		log.Fatal("Environment variable INFURA_URL is not set")
	}
	return infuraURL
}

func GetAccountData(client *ethclient.Client, address string) (string, uint64, error) {
	account := common.HexToAddress(address)

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", 0, err
	}

	balanceInEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))

	balanceInString := balanceInEth.Text('f', 18)

	nonce, err := client.NonceAt(context.Background(), account, nil)
	if err != nil {
		return "", 0, err
	}

	return balanceInString, nonce, nil
}

func GetERC20Balance(client *ethclient.Client, address, tokenAddress string) (string, error) {
	account := common.HexToAddress(address)
	tokenAddr := common.HexToAddress(tokenAddress)

	erc20ABI, err := abi.JSON(strings.NewReader(`
		[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"}]
	`))

	if err != nil {
		return "", err
	}

	data, err := erc20ABI.Pack("balanceOf", account)
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		To:   &tokenAddr,
		Data: data,
	}

	ctx := context.Background()
	output, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "nil", err
	}

	results, err := erc20ABI.Unpack("balanceOf", output)
	if err != nil {
		return "nil", err
	}

	if len(results) == 0 {
		return "nil", errors.New("no balance returned")
	}

	balance, ok := results[0].(*big.Int)
	if !ok {
		return "nil", errors.New("unable to type assert balance to *big.Int")
	}

	balanceInEth := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))

	balanceInString := balanceInEth.Text('f', 18)

	return balanceInString, nil
}
