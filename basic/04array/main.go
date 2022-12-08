package main

import "fmt"

func main() {
	// array 下面两个不是同一个类型
	var arr1 [3]int
	var arr2 [4]int
	arr3 := [3]int{2, 3, 6}        // 初始化
	var arr4 = [4]int{2, 3, 64, 8} // 初始化

	var arr5 = [...]int{2, 3, 64, 8} // 初始化
	arr1 = [3]int{2, 3, 6}
	arr1[0] = 1
	arr2[0] = 1
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	// 遍历
	for i := 0; i < len(arr1); i++ {
		fmt.Println(i)
	}
	for key, value := range arr3 {
		fmt.Println(key, value)
	}

	//数组的比较(类型一样才能比较)
	if arr4 == arr5 {
		fmt.Println("一样")
	}

	// 多维数组
	var arrPlus [3][4]int
	arrPlus[0] = [4]int{1, 2, 3, 4}
	arrPlus[1] = [4]int{1, 2, 3, 5}
	arrPlus[2] = [4]int{1, 2, 3, 3}
	fmt.Println(arrPlus)

}
