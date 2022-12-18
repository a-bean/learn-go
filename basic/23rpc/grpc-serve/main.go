package main

import (
	"context"
	"google.golang.org/grpc"
	p "learn-go/basic/23rpc/grpc"
	"net"
)

type Server struct {
	p.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *p.HelloRequest) (
	*p.HelloReply,
	error) {
	return &p.HelloReply{
		Message: "hello" + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	p.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
