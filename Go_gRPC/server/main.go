package main

import (
	"fmt"
	"log"
	"net"

	//"github.com/plutov/packagemain/00-grpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello there!............")
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("unable to listen on 8080 port: %v", err)
	}

	srv := grpc.NewServer()

	proto.RegisterBlockchainServer(srv, &Server{})
	srv.Serve(listener)

}

type Server struct{}

func (s *Server) AddBlock(context.Context, *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	return new(proto.AddBlockResponse), nil
}

func (s *Server) GetBlockChain(context.Context, *proto.GetBlockRequest) (*proto.GetBlockResponse, error) {
	return new(proto.GetBlockResponse), nil
}
