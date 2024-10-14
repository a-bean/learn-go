package main

import "fmt"

// House 产品
type House struct {
	door   string
	window string
}

// HouseBuilder 建造者
type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) SetDoor(door string) *HouseBuilder {
	b.house.door = door
	return b
}

func (b *HouseBuilder) SetWindow(window string) *HouseBuilder {
	b.house.window = window
	return b
}

func (b *HouseBuilder) Build() House {
	return b.house
}

func main() {
	builder := &HouseBuilder{}
	house := builder.SetDoor("Wooden Door").SetWindow("Glass Window").Build()
	fmt.Println("House:", house)
}
