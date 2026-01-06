package main

import (
	"context"
	"log"
	pb "unary/proto"
)

func sayHello(client pb.HelloServiceClient) {
	log.Println("sayHello is invoked")
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Will",
	})

	if err != nil {
		log.Fatalf("Could not say hello: %v\n", err)
	}

	log.Printf("gRPC response: %s\n", resp.Message)
}
