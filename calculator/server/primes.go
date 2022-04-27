package main

import (
	"log"

	pb "github.com/yoelpater/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Sum Functions was invoked with %v\n", in)

	N := int(in.Number)

	for k := 2; N > 1; {

		if N%k == 0 {

			stream.Send(&pb.PrimesResponse{
				Result: int32(k),
			})
			N = N / k
		} else {
			k = k + 1
		}

	}

	return nil
}

// Primes(*PrimesRequest, CalculatorService_PrimesServer) error
