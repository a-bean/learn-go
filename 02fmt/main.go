package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name := "kobe"
	age := 18
	fmt.Println(name)
	fmt.Println("名称：", name)

	fmt.Printf("名称：%s \r\n年龄: %d", name, age) // 性能不好

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
