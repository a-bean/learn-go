package main

import "fmt"

/*
map的元素无法取址的原因:
	map的扩容跟重整会导致键值的重新分布
	元素还会因为各种原因导致地址变化
*/

func main() {
	// map 是一个key，value的 无序 集合，查询方便O(1)
	var map1 = map[string]string{
		"name": "kobe",
	}
	map1["age"] = "18"
	fmt.Println(map1["name"])

	var mapFloat = map[float64]int{
		1.1: 1,
		1.2: 2,
	}
	fmt.Println(mapFloat)

	//var map2 map[int]int
	// 初始化
	var map3 = map[int]int{}
	var map4 = make(map[int]int, 2) //常用
	//map2[1] = 2          // 会报错,必须先初始化
	//fmt.Println(map2[1]) //panic: assignment to entry in nil map
	fmt.Println(map3)
	fmt.Println(map4)

	// 遍历
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	for key := range map1 {
		fmt.Println(key)
	}
	// 判断某个元素是否存在，
	fmt.Println(map1["curry"]) //不能判断它是不是一个空值

	if data, ok := map1["curry"]; ok { //正确写法
		fmt.Println(data)
	}

	// 删除
	delete(map1, "age")

	// map不是线程安全的，多个go routine进行操作会报错的
}
