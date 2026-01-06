package main

import (
	"context"
	"log"
	helloPb "unary/proto/hello"
)

func sayHello(client helloPb.HelloServiceClient) {
	log.Println("sayHello is invoked")
	resp, err := client.SayHello(context.Background(), &helloPb.HelloRequest{
		Name: "Will",
	})

	if err != nil {
		log.Fatalf("Could not say hello: %v\n", err)
	}

	log.Printf("gRPC response: %s\n", resp.Message)
}
