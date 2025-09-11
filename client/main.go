package main

import (
	"log"

	pb "github.com/Prayag2003/grpc-proto-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"A", "B", "C"},
	}

	callSayHello(client)
	callSayHelloServerStream(client, names)
}
