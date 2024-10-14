package main

import "fmt"

// 状态模式允许对象在其内部状态改变时改变其行为，似乎该对象改变了其类。

// 状态接口
type State interface {
	doAction(context *Context)
}

// 具体状态
type StartState struct{}

func (s *StartState) doAction(context *Context) {
	fmt.Println("Player is in start state")
	context.setState(s)
}

type StopState struct{}

func (s *StopState) doAction(context *Context) {
	fmt.Println("Player is in stop state")
	context.setState(s)
}

// 上下文
type Context struct {
	state State
}

func (c *Context) setState(state State) {
	c.state = state
}

func (c *Context) getState() State {
	return c.state
}

func main() {
	context := &Context{}

	startState := &StartState{}
	startState.doAction(context)

	fmt.Println("Current State:", context.getState())

	stopState := &StopState{}
	stopState.doAction(context)

	fmt.Println("Current State:", context.getState())
}
