package main

import "fmt"

//命令模式:将请求封装成对象，从而使得可以用不同的请求对客户端进行参数化。

// 命令接口
type Command interface {
	Execute()
}

// 具体命令
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// 接收者
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

// 调用者
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.command = c
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}

	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton() // 输出: Light is ON

	remote.SetCommand(lightOff)
	remote.PressButton() // 输出: Light is OFF
}
