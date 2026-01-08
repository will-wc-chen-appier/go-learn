package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	addPb "unary/proto/add"
	helloPb "unary/proto/hello"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type HelloServer struct {
	helloPb.HelloServiceServer
}

type AddServer struct {
	addPb.AddServiceServer
}

func main() {
	grpcAddr := "0.0.0.0:50051"
	httpAddr := "0.0.0.0:8080"

	// Start gRPC server
	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 100,
		}),
		grpc.ReadBufferSize(128*1024),
		grpc.WriteBufferSize(128*1024),
	)
	helloPb.RegisterHelloServiceServer(server, &HelloServer{})
	addPb.RegisterAddServiceServer(server, &AddServer{})
	reflection.Register(server)

	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to listen on: %v\n", err)
		}
		log.Printf("gRPC server listening on: %s\n", grpcAddr)
		if err = server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v\n", err)
		}
	}()

	// grpc-http gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := helloPb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v\n", err)
	}
	go func() {
		log.Printf("HTTP gateway listening on: %s\n", httpAddr)
		if err := http.ListenAndServe(httpAddr, mux); err != nil {
			log.Fatalf("Failed to serve HTTP gateway: %v\n", err)
		} else {
			log.Printf("HTTP listening")
		}
	}()

	// graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	log.Println("shutting down servers...")
	server.GracefulStop()
}
