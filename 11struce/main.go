package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 结构体上定义方法
func (p Person) getName() string {
	return p.name
}

func (p *Person) getName1() string {
	return p.name
}

type Student struct {
	//p      Person // 第一种嵌套方式
	Person // 第二种
	score  int
}

func main() {
	p1 := Person{
		name: "kobe",
		age:  24,
	}
	fmt.Println(p1.getName())
	fmt.Println(p1.name)

	//匿名结构体
	address := struct {
		city string
	}{
		city: "厦门",
	}
	fmt.Println(address)

	// 结构体嵌套
	s := Student{
		Person: Person{
			name: "kobe",
			age:  24,
		},
		score: 100,
	}
	fmt.Println(s)
}
