package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rishabhagrawalzra/gRPCBasicWithGo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
