package main

import (
	"log"
	"net"

	pb "unary/proto"
	addPb "unary/proto/add"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type HelloServer struct {
	pb.HelloServiceServer
}

type AddServer struct {
	addPb.AddServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloServer{})
	addPb.RegisterAddServiceServer(s, &AddServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v\n", err)
	}
}
