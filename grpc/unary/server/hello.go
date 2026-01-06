package main

import (
	"context"
	"log"
	helloPb "unary/proto/hello"
)

func (*HelloServer) SayHello(ctx context.Context, in *helloPb.HelloRequest) (*helloPb.HelloResponse, error) {
	log.Printf("Greet was invoked with %v\n", in)
	return &helloPb.HelloResponse{Message: "Hello " + in.Name}, nil
}
