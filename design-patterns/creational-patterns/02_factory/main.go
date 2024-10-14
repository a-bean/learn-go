package main

import "fmt"

//	普通工厂模式

// Animal 抽象产品
type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof!"
}
func (d Cat) Speak() string {
	return "Woof!"
}

// AnimalFactory 工厂函数
func AnimalFactory(animal string) Animal {
	switch animal {
	case "cat":
		return Cat{}
	case "dog":
		return Dog{}
	default:
		return nil
	}
}

// 抽象工厂模式

// 1. 定义抽象产品接口

// Animal1 接口
type Animal1 interface {
	Speak() string
}

type Plant interface {
	Grow() string
}

// 2. 定义具体产品

// Cat1 具体动物：猫
type Cat1 struct{}

func (c *Cat1) Speak() string {
	return "Meow"
}

// Dog1 具体动物：狗
type Dog1 struct{}

func (d *Dog1) Speak() string {
	return "Woof"
}

// Rose 具体植物：玫瑰
type Rose struct{}

func (r *Rose) Grow() string {
	return "Growing a rose"
}

// Sunflower 具体植物：向日葵
type Sunflower struct{}

func (s *Sunflower) Grow() string {
	return "Growing a sunflower"
}

// 3. 定义抽象工厂接口

// AbstractFactory 接口
type AbstractFactory interface {
	CreateAnimal() Animal1
	CreatePlant() Plant
}

// 4. 定义具体工厂

// AnimalPlantFactory1 具体工厂：工厂1
type AnimalPlantFactory1 struct{}

func (f *AnimalPlantFactory1) CreateAnimal() Animal1 {
	return &Cat{}
}

func (f *AnimalPlantFactory1) CreatePlant() Plant {
	return &Rose{}
}

// AnimalPlantFactory2 具体工厂：工厂2
type AnimalPlantFactory2 struct{}

func (f *AnimalPlantFactory2) CreateAnimal() Animal1 {
	return &Dog{}
}

func (f *AnimalPlantFactory2) CreatePlant() Plant {
	return &Sunflower{}
}

func main() {
	animal := AnimalFactory("dog")
	fmt.Println(animal.Speak()) // 输出: Woof!

	var factory AbstractFactory

	// 抽象工厂模式
	// 使用工厂1
	factory = &AnimalPlantFactory1{}
	animal1 := factory.CreateAnimal()
	plant1 := factory.CreatePlant()

	fmt.Println(animal1.Speak()) // 输出: Meow
	fmt.Println(plant1.Grow())   // 输出: Growing a rose

	// 使用工厂2
	factory = &AnimalPlantFactory2{}
	animal2 := factory.CreateAnimal()
	plant2 := factory.CreatePlant()

	fmt.Println(animal2.Speak()) // 输出: Woof
	fmt.Println(plant2.Grow())   // 输出: Growing a sunflower
}
