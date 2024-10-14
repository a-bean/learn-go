// 装饰器模式: 允许动态地为对象添加行为，而不改变其结构。
package main

import "fmt"

// Coffee 接口
type Coffee interface {
	Cost() int
}

// SimpleCoffee 具体构件
type SimpleCoffee struct{}

func (c SimpleCoffee) Cost() int {
	return 5
}

// MilkDecorator 装饰器
type MilkDecorator struct {
	coffee Coffee
}

func (d MilkDecorator) Cost() int {
	return d.coffee.Cost() + 2
}

type SugarDecorator struct {
	coffee Coffee
}

func (d SugarDecorator) Cost() int {
	return d.coffee.Cost() + 1
}

func main() {
	coffee := SimpleCoffee{}
	fmt.Println("Simple coffee cost:", coffee.Cost())

	milkCoffee := MilkDecorator{coffee: coffee}
	fmt.Println("Milk coffee cost:", milkCoffee.Cost())

	milkSugarCoffee := SugarDecorator{coffee: milkCoffee}
	fmt.Println("Milk & Sugar coffee cost:", milkSugarCoffee.Cost())
}
