package main

import "fmt"

func Add[T int | int32](a, b T) T {
	return a + b
}

type Mymap[KEY int | string, VALUE int | string] map[KEY]VALUE

// Company 结构体
type Company[T string] struct {
	Name string
	CEO  T
}

// Mychannel chan
type Mychannel[T int] chan T

// Company1 嵌套
type Company1[T string, S []T] struct {
	Name string
	Work T
	CEO  S
}

func main() {
	a := Add[int](1, 6)
	fmt.Println(a)

	// map泛型
	m := Mymap[int, int]{1: 2}
	fmt.Println(m)

	// 常见的错误用法
	//1.类型参数不能单独使用
	// type CommonType[T int] T
	type CommonType[T int] []T

	//2. 指针
	//type CommonType1[T *int] []T
	type CommonType1[T interface{ *int }] []T

	// 匿名struct不支持泛型，匿名函数也不支持
	//泛型不支持switch断言
}
