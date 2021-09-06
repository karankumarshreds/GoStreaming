package payment

import (
	"fmt"
	"time"
	"log"
	pb "github.com/karankumarshreds/GoStreaming/protofiles"
)

const (
	streamSteps = 3
)

// keep uppercase to make it accessible by other packages
type PaymentServer struct{}

// implements MakeTransaction from pb (check the compiled code)
func (p *PaymentServer) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	
	log.Printf("Got amount %v : ", in.Amount)
	log.Printf("Got from %v : ", in.From)
	log.Printf("Got for %v : ", in.To)

	// send stream
	for i := 0; i < streamSteps; i++ {
		// send message to the client every 2 seconds 
		time.Sleep(time.Second*2)
		err := stream.Send(&pb.TransactionResponse{
			Step: int32(i+1),
			Status: "Good",
			Description: fmt.Sprintf("Executing step %v", i+1),
		})
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
	}

	log.Println("Successfully streamed the data!")
	return nil

}