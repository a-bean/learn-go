package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		err panic recover（相当于其他函数的try catch）
		1. 开发人员设计一个函数的时候，需要返回一个err值告诉调用者是否成功
		2. go设计者要求必须处理err。（防御编程）

		panicking的执行过程:
		1.发生panic的函数立即停止执行,已求值的defer(在panic之前声明的)继续执行
		2.对于函数调用着而言,调用发生panic跟直接panic类似

		panic和recover一些细节点
		1. recover必须在defer声明的匿名函数中执行
		2. recover只能捕获同一个协程的panic
		3. 当前的goroutine中的panic会被defer中的panic覆盖
		4. 多个defer中的panic执行顺序
		5. 多个调用链中捕获panic,会优先被当前的协程的recover捕获
	*/
	recover1()
	fmt.Println("继续执行")
	deferPanic() // defer2 defer1 然后在panic
}

func deferPanic() {
	defer func() {
		fmt.Println("defer1")
		panic("err1")
	}()

	defer func() {
		fmt.Println("defer2")
		panic("err2")
	}()

	panic("deferPanic")
}

func err1() (int, error) {
	return 1, errors.New("this is an error")
}

/*
1. panic 会导致你的程序退出（不推荐使用）
2. 使用场景：比如一个服务的启动，需要依赖其他服务的启动，这时候其他服务要是没有起来，就可以使用panic
*/
func panic1() {
	panic("this is a panic")
}

/*
recover：这个函数用来捕获panic
*/
func recover1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover,", r)
		}
	}()
	var names map[string]string
	names["name1"] = "kobe"
}
