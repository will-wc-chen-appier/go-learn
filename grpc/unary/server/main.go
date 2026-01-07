package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	addPb "unary/proto/add"
	helloPb "unary/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HelloServer struct {
	helloPb.HelloServiceServer
}

type AddServer struct {
	addPb.AddServiceServer
}

func main() {
	server := grpc.NewServer()
	helloPb.RegisterHelloServiceServer(server, &HelloServer{})
	addPb.RegisterAddServiceServer(server, &AddServer{})
	defer server.GracefulStop()

	reflection.Register(server)

	go func() {
		var addr string = "0.0.0.0:50051"
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Failed to listen on: %v\n", err)
		}
		if err = server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v\n", err)
		}
		log.Printf("gRPC server listening on: %v\n", addr)
	}()

	// graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	log.Println("shutting down gRPC server...")
}
