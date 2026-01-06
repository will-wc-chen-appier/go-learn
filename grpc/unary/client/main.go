package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	addPb "unary/proto/add"
	helloPb "unary/proto/hello"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := helloPb.NewHelloServiceClient(conn)
	sayHello(client)
	addClient := addPb.NewAddServiceClient(conn)
	addTwoNumbers(addClient, 1, 2)
}
