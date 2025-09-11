package main

import (
	"log"
	"time"

	pb "github.com/Prayag2003/grpc-proto-go/proto"
)

func (s *helloServer) SayHelloFromServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloFromServerStreamingServer) error {
	log.Printf("Got request for names: %v", req)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
