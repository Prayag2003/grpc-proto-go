package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Prayag2003/grpc-proto-go/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloFromClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished!")
	if err != nil {
		log.Fatalf("Error while receiving the response %v", err)
	}
	log.Printf("%v", res.Messages)
}
