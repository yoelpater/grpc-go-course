package main

import (
	"log"
	"os"
	"strconv"

	pb "github.com/yoelpater/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5001"

func main() {

	if len(os.Args) < 2 {
		log.Fatal("required 1 argument\n")
		return
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//sum(c)

	num := os.Args[1]

	num1, err := strconv.Atoi(num)

	if err != nil {
		log.Fatalf("failed to parse number %v\n", err)
	}

	doPrimes(c, int32(num1))

}
