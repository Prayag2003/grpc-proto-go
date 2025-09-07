package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
