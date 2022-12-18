package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	ipc "learn-go/basic/23rpc/proto"
)

func main() {
	req := ipc.HelloRequest{
		Name:    "kobe24",
		Age:     24,
		Courses: []string{"go", "gin", "微服务"},
	}
	rsp, _ := proto.Marshal(&req) // 序列化
	fmt.Println(rsp)
	fmt.Println(string(rsp))

	newReq := ipc.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq) // 反序列化
	fmt.Println(newReq.Name)
	fmt.Println(newReq.Age)
	fmt.Println(newReq.Courses)
}
