package main

import (
	"net"
	"net/rpc"
)

// HelloService 注册接口
type HelloService struct{}

func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	// 1. 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")
	// 2. 注册处理逻辑handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	// 3. 启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn) // golang默认的协议 gob 序列化
	// rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // json 序列化
}
