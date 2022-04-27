package main

import (
	"context"
	"log"

	pb "github.com/yoelpater/grpc-go-course/calculator/proto"
)

func sum(c pb.CalculatorServiceClient) {
	log.Println("calculator sum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum: %d", res.Result)
}
