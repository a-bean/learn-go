package main

import "fmt"

// 服务定位器模式通过在运行时查找和提供服务来简化服务访问，通常用于解决复杂依赖关系。

// 服务接口
type Service interface {
	GetName() string
}

// 具体服务
type ConcreteServiceA struct{}

func (s *ConcreteServiceA) GetName() string {
	return "Service A"
}

type ConcreteServiceB struct{}

func (s *ConcreteServiceB) GetName() string {
	return "Service B"
}

// 服务定位器
type ServiceLocator struct {
	services map[string]Service
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		services: make(map[string]Service),
	}
}

func (locator *ServiceLocator) GetService(serviceName string) Service {
	if service, exists := locator.services[serviceName]; exists {
		return service
	}
	var service Service
	if serviceName == "ServiceA" {
		service = &ConcreteServiceA{}
	} else if serviceName == "ServiceB" {
		service = &ConcreteServiceB{}
	}
	locator.services[serviceName] = service
	return service
}

func main() {
	locator := NewServiceLocator()

	serviceA := locator.GetService("ServiceA")
	fmt.Println(serviceA.GetName()) // 输出: Service A

	serviceB := locator.GetService("ServiceB")
	fmt.Println(serviceB.GetName()) // 输出: Service B
}
