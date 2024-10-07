package server

import (
	"context"
	"errors"
	"io"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github/serlenario/go-gRPC-ethereum/internal/ethereum"
	"github/serlenario/go-gRPC-ethereum/internal/proto"
	"github/serlenario/go-gRPC-ethereum/internal/validator"
)

type AccountServer struct {
	proto.UnimplementedAccountServiceServer
}

func (s *AccountServer) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	ethereumAddress := req.GetEthereumAddress()
	cryptoSignature := req.GetCryptoSignature()

	isValid, err := validation.ValidateSignature(ethereumAddress, cryptoSignature)
	if err != nil || !isValid {
		return nil, errors.New("invalid signature")
	}

	client, err := ethclient.Dial(ethereum.InfuraURL())
	if err != nil {
		return nil, err
	}
	defer client.Close()

	gasTokenBalance, walletNonce, err := ethereum.GetAccountData(client, ethereumAddress)
	if err != nil {
		return nil, err
	}

	return &proto.GetAccountResponse{
		GasTokenBalance: gasTokenBalance,
		WalletNonce:     walletNonce,
	}, nil
}

func (s *AccountServer) GetAccounts(stream proto.AccountService_GetAccountsServer) error {
	client, err := ethclient.Dial(ethereum.InfuraURL())
	if err != nil {
		return err
	}
	defer client.Close()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		erc20TokenAddress := req.GetErc20TokenAddress()
		ethereumAddresses := req.GetEthereumAddresses()

		for _, address := range ethereumAddresses {
			balance, err := ethereum.GetERC20Balance(client, address, erc20TokenAddress)
			if err != nil {
				log.Printf("Error getting balance for %s: %v", address, err)
				continue
			}

			resp := &proto.GetAccountsResponse{
				EthereumAddress: address,
				Erc20Balance:    balance,
			}

			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
}
