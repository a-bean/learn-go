// 外观模式: 为复杂子系统提供一个简化的接口。
package main

import "fmt"

// 子系统
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU is starting")
}

type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory is loading")
}

type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("Hard drive is reading data")
}

// 外观
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (f *ComputerFacade) Start() {
	f.cpu.Start()
	f.memory.Load()
	f.hardDrive.Read()
}

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
