package main

import (
	"context"
	"fmt"
	"log"
	"os"
	pb "streaming_client/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50052"
	defaultName = "gRPC"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.SayHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("sending %v into the stream\n", i)
		stream.Send(&pb.HelloRequest{Name: name})
		time.Sleep(100 * time.Millisecond)
	}
}
