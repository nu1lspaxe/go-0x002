package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "streaming_bidirectional/proto"

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
	stream, err := c.SayHello(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Blocking channel
	waitc := make(chan struct{})

	// Send feeds to the stream
	go func() {
		for i := 1; i <= 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("error while sending feed: %v", err)
			}
			time.Sleep(time.Second)
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("failed to close stream: %v", err)
		}
	}()

	// Receive feeds from the stream
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("failed to recieve: %v", err)
				close(waitc)
				return
			}
			fmt.Printf("New feed recieved : %v\n", msg.GetMessage())
		}
	}()

	<-waitc
}
