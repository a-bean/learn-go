package main

import "fmt"

func main() {
	/*
			go中的函数是一等公民
				1. 函数本身可以作为变量
				2. 匿名函数，闭包。
				3. 函数可以满足接口
		函数传递参数的时候，是值传递（go中全部是值传递）
	*/
	fmt.Println(add(1, 3))
	add3(1, 2, 3, 6)
}

func add(a, b int) int {
	return a + b
}

// 多值返回
func add1(a, b int) (int, error) {
	return a + b, nil
}

// 返回值变量
func add2(a, b int) (sum int, err error) {
	sum = a + b
	err = nil
	return
}

// 可变参数
func add3(a ...int) int { // a是一个slice
	fmt.Println(a)
	return 0
}
func add4(b string, a ...int) int { // a是一个slice
	fmt.Println(b, a)
	return 0
}

// 一等公民的特性
func add5(fn func(int) int) int {
	return fn(5)
}

func add6(a int) func() {
	return func() {
		fmt.Println(a)
	}
}

func add7(fn func(int) int) func(int) int {
	return fn
}

var add8 = func(a int) int {
	return a
}
