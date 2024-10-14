package main

// 策略模式: 定义了一系列算法，将每个算法封装起来，使它们可以互换。

import "fmt"

// Strategy 策略接口
type Strategy interface {
	Execute(int, int) int
}

// Add 具体策略
type Add struct{}
type Subtract struct{}

func (a Add) Execute(x, y int) int {
	return x + y
}

func (s Subtract) Execute(x, y int) int {
	return x - y
}

// Context 上下文
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExecuteStrategy(x, y int) int {
	return c.strategy.Execute(x, y)
}

func main() {
	context := &Context{}

	context.SetStrategy(Add{})
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5)) // 输出: 15

	context.SetStrategy(Subtract{})
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5)) // 输出: 5
}
