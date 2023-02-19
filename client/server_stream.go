package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rishabhagrawalzra/gRPCBasicWithGo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send the names: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		} else {
			log.Println(msg)
		}
	}
	log.Printf("Straming Finished")
}
