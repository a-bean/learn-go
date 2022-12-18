package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		return
	}
	var reply string
	err1 := client.Call("HelloService.Hello", "kobe", &reply)
	if err1 != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
