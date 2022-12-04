package main

import (
	"fmt"
	"strconv"
)

// 全局变量
var (
	name1 string
	age1  = 18
)

func main() {
	// 第一种
	var name string
	name = "eee"
	fmt.Println(name)

	// 第二种
	var age = 18
	fmt.Println(age)

	// 第三种
	sex := "男"
	fmt.Println(sex)

	// 多变量
	var user1, user2, user3 = "kobe", "curry", "klay"
	fmt.Println(user1, user2, user3)

	// 常量
	const PI = 5
	const (
		PI1 = 1
		PI2 = 2
		PI3
		PI4 = "ab"
		PI5 // 沿用上面的值
	)
	fmt.Println(PI, PI1, PI2, PI3, PI4, PI5)

	// iota 特殊常量，可以被编译器修改
	const (
		ERR1 = iota
		ERR2
		ERR3 = "h"
		ERR4
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR4)

	// 匿名变量
	var _ int

	// byte -> uint8 主要适用于存放字符
	var c byte
	c = 'c'
	fmt.Println(c)
	fmt.Printf("c=%c", c)

	// rune -> int32 也是字符
	var r rune
	r = '字'
	fmt.Println(r)
	fmt.Printf("r=%c", r)

	// 类型转换
	temp := 1.0
	temp1 := int(temp)
	temp2 := int8(temp)
	temp3 := uint8(temp)
	temp4 := float64(temp3)
	fmt.Println(temp1, temp2, temp3, temp4)

	// string -> int
	var istr = "12"
	mayInt, err := strconv.Atoi(istr)
	if err == nil {
		fmt.Println(mayInt)
	}

	// int -> string
	myi := 12
	mys := strconv.Itoa(myi)
	fmt.Println(mys)

}
