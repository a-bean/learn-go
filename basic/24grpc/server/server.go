package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	p "learn-go/basic/23rpc/grpc"
	"net"
)

type Server struct {
	p.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *p.HelloRequest) (
	*p.HelloReply,
	error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata err")
	}

	for key, val := range md {
		fmt.Println(key, val)
	}

	return &p.HelloReply{
		Message: "hello! " + request.Name,
	}, nil
}

func interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp any, err error) {
	fmt.Println("收到一个新的请求")
	return handler(ctx, req)
}

func main() {

	opt := grpc.UnaryInterceptor(interceptor) // 拦截器
	g := grpc.NewServer(opt)

	p.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}

}
