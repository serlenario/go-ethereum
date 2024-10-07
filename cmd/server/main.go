package main

import (
	"log"
	"net"

	"github/serlenario/go-gRPC-ethereum/internal/proto"
	"github/serlenario/go-gRPC-ethereum/internal/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterAccountServiceServer(s, &server.AccountServer{})

	log.Println("Server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
