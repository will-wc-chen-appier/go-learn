package main

import (
	"context"
	"log"
	addPb "unary/proto/add"
)

func (*AddServer) Add(c context.Context, req *addPb.AddRequest) (*addPb.AddResponse, error) {
	log.Printf("Add invoked by A: %v and B: %v\n", req.A, req.B)
	sum := req.A + req.B
	log.Printf("Returned value: %v\n", sum)
	return &addPb.AddResponse{Result: sum}, nil
}
