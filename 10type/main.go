package main

import "fmt"

// type 关键字：定义结构体，接口，类型别名，类型定义，类型判断

type myInt = int // 别名
type myInt1 int  // 自定义类型（可以用来拓展方法）

func main() {
	var a myInt
	var b myInt1
	var c int = 1

	fmt.Println(a + c)
	//fmt.Println(b + c) //报错，因为他们不是同一个类型

	fmt.Printf("myint1 %T \r\n", a)
	fmt.Printf("myint2 %T \r\n", b)

	// switch
	var i any = "abc"
	switch i.(type) {
	case myInt:
		fmt.Println("abc")
	}
}
