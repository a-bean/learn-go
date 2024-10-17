package main

import "fmt"

// 享元模式减少创建对象的数量，以减少内存使用并提高性能。在需要大量细粒度对象时，这种模式特别有用。

// Flyweight 接口
type Shape interface {
	Draw()
}

// 具体的 Flyweight
type Circle struct {
	color  string
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing Circle of color %s with radius %d\n", c.color, c.radius)
}

// Flyweight 工厂
type ShapeFactory struct {
	circleMap map[string]*Circle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]*Circle),
	}
}

func (f *ShapeFactory) GetCircle(color string) *Circle {
	if circle, exists := f.circleMap[color]; exists {
		return circle
	}
	newCircle := &Circle{color: color}
	f.circleMap[color] = newCircle
	return newCircle
}

func main() {
	factory := NewShapeFactory()

	redCircle := factory.GetCircle("red")
	redCircle.SetRadius(5)
	redCircle.Draw() // 输出: Drawing Circle of color red with radius 5

	anotherRedCircle := factory.GetCircle("red")
	anotherRedCircle.SetRadius(10)
	anotherRedCircle.Draw() // 输出: Drawing Circle of color red with radius 10

	// 注意，这两个 Circle 对象是同一个共享对象
	fmt.Println(redCircle == anotherRedCircle) // 输出: true
}
