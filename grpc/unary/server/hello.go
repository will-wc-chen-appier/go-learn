package main

import (
	"context"
	"log"
	pb "unary/proto"
)

func (*Server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Greet was invoked with %v/n", in)
	return &pb.HelloResponse{Message: "Hello" + in.Name}, nil
}
