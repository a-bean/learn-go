package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	p "learn-go/basic/23rpc/proto-stream"
	"sync"
	"time"
)

func main() {
	// 服务端流模式
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := p.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &p.SteamReqData{Data: "kobe"})
	for {
		a, err := res.Recv()
		if err != nil {
			break
		}
		fmt.Println(a)
	}

	// 客户端流模式
	putStr, _ := c.PutStream(context.Background())
	i := 0
	for {
		putStr.Send(&p.SteamReqData{Data: "curry"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			a, err := allStr.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(a)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			allStr.Send(&p.SteamReqData{Data: fmt.Sprintf("%v", time.Now().Unix())})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
