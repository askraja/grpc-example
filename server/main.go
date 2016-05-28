package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
}
func (s *server) SayNothing(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*NothingReply, error) {
}
func (s *server) RememberSomething(ctx context.Context, in *RememberRequest, opts ...grpc.CallOption) (*NothingReply, error) {
}
func (s *server) AskSomething(ctx context.Context, in *RememberRequest, opts ...grpc.CallOption) (*AskResponse, error) {
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
