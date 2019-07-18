package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ding-harrison/SmallSystem/pkg/protobuf"
)

type server struct{}

const (
	port = ":50051"
)

func (s *server) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", r.Name)
	return &pb.HelloReply{Message: "Hello " + r.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
