package main

import "fmt"

type Person struct {
	Age int
}

func main() {
	// 定义
	var p1 *string
	name := "kobe"
	p1 = &name

	fmt.Println(p1)
	fmt.Println(*p1)

	// 初始化
	var p4 = &Person{}
	var p5 = new(Person) //推荐
	var p6 = new(int)    //推荐
	p2 := &Person{
		Age: 24,
	}

	fmt.Println(p4)
	fmt.Println(p5)
	fmt.Println(p6)

	// go中 结构体和结构体指针都能通过点的形式取值
	fmt.Println(p2.Age)
	fmt.Println((*p2).Age)

	// go中的指针不能参与运算 (只有在unsafe包中可以)
	age := 12
	p3 := &age
	// p3++ 报错
	fmt.Println(*p3)

	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)

	/*
		 nil
		不同类型值的零值
		bool: false
		int,float: 0
		string:""
		slice,map,pointer,channel,interface,func: nil
	*/

}

func swap(a, b *int) {
	*a, *b = *b, *a
}
