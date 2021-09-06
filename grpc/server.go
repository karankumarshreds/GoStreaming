package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/karankumarshreds/GoStreaming/protofiles"
	payment "github.com/karankumarshreds/GoStreaming/payment"
)

const (
	port = ":8000"
)

func main() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", port)	
	if err != nil {
		log.Fatalf("Error while listening %v", err)
	}

	pb.RegisterMoneyTransactionServer(grpcServer, )
}
