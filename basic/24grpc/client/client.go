package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	p "learn-go/basic/23rpc/grpc"
	"time"
)

func interceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("收到一个新的请求111")
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("耗时：%s\n", time.Since(start))
	return err
}

func main() {
	opts := grpc.WithUnaryInterceptor(interceptor) //拦截器
	conn, err := grpc.Dial("0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()), opts)
	if err != nil {
		//panic(err)
		// 解析错误（拿到code跟message）
		st, ok := status.FromError(err)
		if !ok {
			panic("解析err失败")

		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	defer conn.Close()
	c := p.NewGreeterClient(conn)

	// metadata
	md := metadata.New(map[string]string{
		"name": "kobe",
		"age":  "24",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 设置超时时间
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	r, err1 := c.SayHello(ctx, &p.HelloRequest{Name: "kobe"})
	if err1 != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
