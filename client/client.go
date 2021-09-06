package main

import (
	"io"
	"log"
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	pb "github.com/karankumarshreds/GoStreaming/protofiles"
)

const grpcAddr = "localhost:8000"

func main() {
	
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection %v", err)
	}

	// client instance
	c := pb.NewMoneyTransactionClient(conn)

	// invoke remote function
	stream, err2 := c.MakeTransaction(
		context.Background(),
		&pb.TransactionRequest{
			Amount: float32(1200.00),
			From: "John",
			To: "Alice",
		},
	)
	if err2 != nil {
		log.Fatalf("Error invoking remote transaction %v", err2)
	}

	// listen to stream of messages 
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			return // assuming no messages left 
		}
		if err != nil {
			log.Fatalf("Error while receiving message %f", err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}

}