package main

import "fmt"

// 备忘录模式用于在不破坏封装的前提下捕获和恢复对象的内部状态。

// 备忘录
type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

// 原发器
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) SaveStateToMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) GetStateFromMemento(m *Memento) {
	o.state = m.GetState()
}

// 管理者
type CareTaker struct {
	mementoList []*Memento
}

func (c *CareTaker) Add(m *Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) Get(index int) *Memento {
	return c.mementoList[index]
}

func main() {
	originator := &Originator{}
	careTaker := &CareTaker{}

	originator.SetState("State #1")
	originator.SetState("State #2")
	careTaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #3")
	careTaker.Add(originator.SaveStateToMemento())

	originator.SetState("State #4")

	fmt.Println("Current State:", originator.state)

	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Println("First saved State:", originator.state)

	originator.GetStateFromMemento(careTaker.Get(1))
	fmt.Println("Second saved State:", originator.state)
}
