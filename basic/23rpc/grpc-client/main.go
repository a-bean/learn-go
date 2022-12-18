package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	p "learn-go/basic/23rpc/grpc"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := p.NewGreeterClient(conn)
	r, err1 := c.SayHello(context.Background(), &p.HelloRequest{Name: "kobe"})
	if err1 != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
