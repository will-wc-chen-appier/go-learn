package main

import (
	"context"
	"log"
	addPb "unary/proto/add"
)

func addTwoNumbers(client addPb.AddServiceClient, a float64, b float64) {
	log.Println("addTwoNumbers is invoked with %v and %v on the client side\n", a, b)
	res, err := client.Add(context.Background(), &addPb.AddRequest{
		A: a,
		B: b,
	})
	if err != nil {
		log.Fatalf("Could not add two numbers: %v\n", err)
	}

	log.Println("the sum of %v and %v is %v\n", a, b, res.Result)
}
