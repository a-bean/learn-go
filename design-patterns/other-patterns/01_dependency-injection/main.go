package main

import "fmt"

// 依赖注入模式通过将对象的依赖关系传递给对象来解耦组件，使得依赖关系可以在运行时动态更改。

// 服务接口
type Service interface {
	Serve()
}

// 具体服务
type ConcreteService struct{}

func (s *ConcreteService) Serve() {
	fmt.Println("Service is serving...")
}

// 注入服务的结构
type Client struct {
	service Service
}

func NewClient(service Service) *Client {
	return &Client{service: service}
}

func (c *Client) DoSomething() {
	c.service.Serve()
}

func main() {
	service := &ConcreteService{}
	client := NewClient(service)
	client.DoSomething() // 输出: Service is serving...
}
