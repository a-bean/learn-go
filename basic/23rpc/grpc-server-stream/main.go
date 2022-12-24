package main

import (
	"fmt"
	"google.golang.org/grpc"
	p "learn-go/basic/23rpc/proto-stream"
	"net"
	"sync"
	"time"
)

const PORT = ":50052"

type server struct {
	p.UnimplementedGreeterServer
}

func (s *server) GetStream(req *p.SteamReqData, res p.Greeter_GetStreamServer) error {
	i := 0
	for {
		res.Send(&p.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}
	return nil
}

func (s *server) PutStream(cliStr p.Greeter_PutStreamServer) error {
	for {
		a, err := cliStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a)
	}
	return nil
}
func (s *server) AllStream(allStr p.Greeter_AllStreamServer) error {
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
			allStr.Send(&p.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	g := grpc.NewServer()
	p.RegisterGreeterServer(g, &server{})
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
