package main

import (
	"context"
	"log"

	pb "github.com/yoelpater/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum Functions was invoked with %d + %d\n", in.FirstNumber, in.SecondNumber)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
