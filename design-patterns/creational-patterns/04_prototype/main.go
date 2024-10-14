package main

import "fmt"

// 原型模式: 用于复制现有的对象而不是创建新的实例。

// Cloneable 原型接口
type Cloneable interface {
	Clone() Cloneable
}

// Person 具体原型
type Person struct {
	name string
	age  int
}

func (p *Person) Clone() Cloneable {
	clone := *p
	return &clone
}

func main() {
	person1 := &Person{name: "John", age: 30}
	person2 := person1.Clone().(*Person)

	person2.name = "Doe"
	fmt.Println(person1.name, person1.age) // 输出: John 30
	fmt.Println(person2.name, person2.age) // 输出: Doe 30
}
