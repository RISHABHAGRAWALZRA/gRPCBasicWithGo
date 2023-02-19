package main

import (
	"context"

	pb "github.com/rishabhagrawalzra/gRPCBasicWithGo/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello there",
	}, nil
}
