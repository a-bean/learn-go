package main

import "fmt"

// 中介者模式定义了一个对象，该对象封装了一组对象之间的交互方式。
//	通过使用中介者，对象不需要相互显式引用，从而实现松耦合。

// 中介者接口
type Mediator interface {
	Send(message string, colleague Colleague)
}

// 抽象同事类
type Colleague interface {
	Send(message string)
	Receive(message string)
}

// 具体中介者
type ConcreteMediator struct {
	colleague1 *ConcreteColleague1
	colleague2 *ConcreteColleague2
}

func (m *ConcreteMediator) Send(message string, colleague Colleague) {
	if colleague == m.colleague1 {
		m.colleague2.Receive(message)
	} else {
		m.colleague1.Receive(message)
	}
}

// 具体同事类
type ConcreteColleague1 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) Send(message string) {
	fmt.Println("Colleague1 sends message:", message)
	c.mediator.Send(message, c)
}

func (c *ConcreteColleague1) Receive(message string) {
	fmt.Println("Colleague1 received message:", message)
}

type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague2) Send(message string) {
	fmt.Println("Colleague2 sends message:", message)
	c.mediator.Send(message, c)
}

func (c *ConcreteColleague2) Receive(message string) {
	fmt.Println("Colleague2 received message:", message)
}

func main() {
	mediator := &ConcreteMediator{}

	colleague1 := &ConcreteColleague1{mediator: mediator}
	colleague2 := &ConcreteColleague2{mediator: mediator}

	mediator.colleague1 = colleague1
	mediator.colleague2 = colleague2

	colleague1.Send("Hello, Colleague2!")
	colleague2.Send("Hi, Colleague1!")
}
