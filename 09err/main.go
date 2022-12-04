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
	*/
	recover1()
	fmt.Println("继续执行")
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
