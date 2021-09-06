package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	pb "github.com/karankumarshreds/GoStreaming/protofiles"
	"github.com/karankumarshreds/GoStreaming/payment"
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

	// register payment service as the money transaction server 
	paymentService := payment.PaymentServer{}
	pb.RegisterMoneyTransactionServer(grpcServer, &paymentService)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error while serving %v", err)
	}
	log.Println("Listening on port", port)

}
