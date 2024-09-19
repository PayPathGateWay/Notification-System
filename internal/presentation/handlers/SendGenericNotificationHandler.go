package handlers

//import (
//	"Notification-System/internal/gRPC/__" // Replace with actual package path
//	"context"
//	"fmt"
//	"google.golang.org/grpc"
//	"log"
//	"net"e
//)
//
//type server struct {
//	___.UnimplementedGreeterServer
//}
//
//func (s *server) SayHelloAgain(
//	ctx context.Context, in *___.HelloRequest) (*___.HelloReply, error) {
//	fmt.Println(in)
//	// send the data to be processed and await to the send
//	if in.GetName() == "Osama" {
//		fmt.Println("I am processing this request")
//	}
//
//	return &___.HelloReply{Message: "Hello again " + in.GetName()}, nil
//}
//
//func StartThegRPCServer() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	___.RegisterGreeterServer(grpcServer, &server{})
//
//	log.Println("gRPC server is listening on port http://localhost:50051/")
//	if err := grpcServer.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
