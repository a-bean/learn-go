package main

import (
	"container/list"
	"fmt"
)

func main() {
	var list1 list.List
	list1.PushBack("kobe")
	list1.PushBack("kobe")
	list1.PushBack("kobe")
	fmt.Println(list1)

	// 遍历
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
