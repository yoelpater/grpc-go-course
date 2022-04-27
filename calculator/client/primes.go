package main

import (
	"context"
	"io"
	"log"

	pb "github.com/yoelpater/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient, number int32) {
	log.Println("doPrimes was invoked")

	// res, err := c.Primes((context.Background(), &pb.SumRequest{
	// 	FirstNumber:  3,
	// 	SecondNumber: 10,
	// }))

	req := &pb.PrimesRequest{
		Number: number,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Primes: %d", msg.Result)
	}
}
