package main

import (
	"io"
	"log"

	pb "github.com/Prayag2003/grpc-proto-go/proto"
)

func (s *helloServer) SayHelloFromClientStreaming(stream pb.GreetService_SayHelloFromClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Request with names :%v", req.Name)
		messages = append(messages, "Hello ", req.Name)
	}
}
