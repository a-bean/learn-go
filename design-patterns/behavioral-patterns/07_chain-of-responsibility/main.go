package main

import "fmt"

// 责任链模式允许多个对象有机会处理请求，从而避免请求的发送者和接收者之间的耦合。这些对象通过形成一条链来处理请求。

// 处理器接口
type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

// 基础处理器
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) HandleRequest(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// 具体处理器
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled request A")
	} else {
		h.BaseHandler.HandleRequest(request)
	}
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled request B")
	} else {
		h.BaseHandler.HandleRequest(request)
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB)

	handlerA.HandleRequest("A")
	handlerA.HandleRequest("B")
}
