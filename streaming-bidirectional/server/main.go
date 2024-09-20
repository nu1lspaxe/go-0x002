package main

import (
	"fmt"
	"io"
	"log"
	"net"
	pb "streaming_bidirectional/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type ChatServer struct {
	pb.UnimplementedGreeterServer
}

func (c *ChatServer) SayHello(srv pb.Greeter_SayHelloServer) error {
	for {
		msg, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("could not recieve from srv : %v", err)
			return err
		}

		fmt.Println("sending new feed...", msg.GetName())
		srv.Send(&pb.HelloReply{Message: "hello"})
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &ChatServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
