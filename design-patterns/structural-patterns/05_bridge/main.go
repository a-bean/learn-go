package main

import "fmt"

// 桥接模式将抽象部分与它的实现部分分离，使它们都可以独立地变化。

// Device 实现接口
type Device interface {
	On()
	Off()
}

// TV 具体实现
type TV struct{}

func (t *TV) On() {
	fmt.Println("TV is ON")
}

func (t *TV) Off() {
	fmt.Println("TV is OFF")
}

// RemoteControl 抽象类
type RemoteControl struct {
	device Device
}

func (r *RemoteControl) TurnOn() {
	r.device.On()
}

func (r *RemoteControl) TurnOff() {
	r.device.Off()
}

func main() {
	tv := &TV{}
	remote := &RemoteControl{device: tv}

	remote.TurnOn()  // 输出: TV is ON
	remote.TurnOff() // 输出: TV is OFF
}
