package main

import (
	"fmt"
)

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

	// 闭包
	nextFn := add9()
	fmt.Println(nextFn())
	fmt.Println(nextFn())
	fmt.Println(nextFn())
	fmt.Println(nextFn())

	add10()
	add11([]int{1, 2, 3})

	var s []int = []int{1, 2, 3}
	setSlice(s)
	fmt.Println(s)
	fmt.Println(deferDemo(1)) // 3

}

func setSlice(s []int) {
	s[0] = 4
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

func add11(a []int) {
	fmt.Println(a)
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

// 闭包
func add9() func() int {
	local := 0 // 一个函数中，访问另外一个函数的局部变量是不行的
	return func() int {
		local += 1
		return local
	}
}

// defer 在return之前执行。用来最后来释放资源等功能
func add10() {
	defer fmt.Println("defer1") // 最后执行
	defer fmt.Println("defer2") // 3
	defer fmt.Println("defer3") // 2
	fmt.Println("add")          // 最先执行
}

func deferDemo(a int) (r int) {
	defer func() {
		r += a
	}()
	return a + a
}
