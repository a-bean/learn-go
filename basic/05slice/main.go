package main

import "fmt"

func main() {
	var slice1 []int
	// slice1[0] = 2 不能怎么做,因为空间还没分配好
	slice1 = append(slice1, 1)
	fmt.Println(slice1[0])
	fmt.Println(slice1)

	// 初始化
	arr := [4]int{112, 3, 6, 4}

	slice2 := arr[0:2]
	fmt.Println(slice2)

	slice3 := []int{112, 3, 6, 4}
	fmt.Println(slice3)

	slice4 := make([]int, 3, 4)
	slice4[0] = 2
	fmt.Println(slice4)

	// 切片的字面量表达式
	s1 := []string{"a", "b", "c"}
	s2 := []string{0: "a", 1: "b", 2: "c"}
	s3 := []string{2: "a", 1: "b", 0: "c"}
	s4 := []string{2: "a", 0: "b", "c"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// 访问 单个和多个
	fmt.Println(slice3[0], slice3[0:3], slice3[:2], slice3[1:], slice3[:])

	// ...
	slice4 = append(slice4, slice3...)
	fmt.Println("slice4", slice4)

	// 删除元素
	slice4 = append(slice4[:1], slice4[2:]...)
	fmt.Println(slice4)

	// 复制slice
	copySlice := slice4 //两个是同一个引用，一个变 另外一个也会变
	fmt.Println(copySlice)
	slice4[1] = 100
	fmt.Println(copySlice)

	// copy不会自动扩容
	var copySlice1 []int
	var copySlice2 = make([]int, 3)
	copy(copySlice1, slice4)
	copy(copySlice2, slice4)
	fmt.Println(copySlice1) // [] 这个打印空  因为copy不会自动扩容
	fmt.Println(copySlice2) //[2 100 112]

	// slice 原理 本质是一个结构体
	// 1. go的slice在函数参数传递的时候是值传递还是引用传递：值传递，效果又呈现引用传递的效果（不完全是）
	slice5 := []int{1, 3, 65}
	print(slice5)
	fmt.Println(slice5)

	aa := []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(aa) //[0 77 88 0 44 55 66]

	// 二维切片初始化
	var slice6 [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(slice6)
}

func print(data []int) {
	data[0] = 100          // 确实改变了外面的变量
	data = append(data, 3) // 改变不了外面的变量，因为扩容了
	//data = append(data[:2], 3) // 改变了外面的变量,因为没有扩容
	fmt.Println(data)
}
