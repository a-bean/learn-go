package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name := "kobe"
	age := 18
	//  Println 直接打印到标准输出，自动加换行
	fmt.Println(name)
	fmt.Println("名称：", name)

	// Printf 格式化打印到标准输出，无自动换行
	fmt.Printf("名称：%s \r\n年龄: %d", name, age) // 性能不好

	// Sprintf 格式化后返回字符串，不直接打印
	msg := fmt.Sprintf("名称：%s \r\n年龄: %d，%v", name, age, age) //首推
	fmt.Println(msg)

	// 字符串拼接 性能nice
	var builder strings.Builder
	builder.WriteString("名称：")
	builder.WriteString(name)
	builder.WriteString("年龄：")
	builder.WriteString(strconv.Itoa(age))
	res := builder.String()
	fmt.Println("res:", res)
}
