package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/AshalIbrahim/userService/proto/userpb"
)

func main() {
	InitDB()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userServer{})

	log.Println("gRPC server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
