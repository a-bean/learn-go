package main

// 空对象模式提供一个无操作的对象来替代 nil，避免显式的 nil 检查。
import "fmt"

// 接口
type Animal interface {
	Speak() string
}

// 具体实现
type Dog struct{}

func (d *Dog) Speak() string {
	return "Woof!"
}

// 空对象
type NullAnimal struct{}

func (n *NullAnimal) Speak() string {
	return "No sound."
}

// 工厂函数
func GetAnimal(isNull bool) Animal {
	if isNull {
		return &NullAnimal{}
	}
	return &Dog{}
}

func main() {
	animal1 := GetAnimal(false)
	fmt.Println(animal1.Speak()) // 输出: Woof!

	animal2 := GetAnimal(true)
	fmt.Println(animal2.Speak()) // 输出: No sound.
}
