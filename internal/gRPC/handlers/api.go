package handlers

import (
	"Notification-System/internal/gRPC/__"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	___.UnimplementedGreeterServer
}

func (s *server) SayHelloAgain(ctx context.Context, in *___.HelloRequest) (*___.HelloReply, error) {
	return &___.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

func StartThegRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	___.RegisterGreeterServer(grpcServer, &server{})

	log.Println("gRPC server is listening on port http://localhost:50051/")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
