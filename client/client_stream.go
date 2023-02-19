package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rishabhagrawalzra/gRPCBasicWithGo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send the names: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending: %v", err)
		}
		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}
	log.Printf("%v", res.Messages)
	log.Printf("Client Streaming Finshed")
}
