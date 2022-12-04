package main

import (
	"fmt"
	"strings"
)

func main() {
	// 长度
	name := "kobe"
	bytes := []byte(name)
	byter := []rune(name)
	fmt.Println(len(name), bytes, byter)
	fmt.Println(len(name), len(bytes), len(byter))

	// 转义符
	curseName := "faf\"efa\""
	curseName1 := `faf"efa"`
	fmt.Println(curseName)
	fmt.Println(curseName1)

	// 字符串比较
	a := "hello"
	b := "good"
	fmt.Println(a == b)
	fmt.Println(a > b) // 比较阿斯克码

	// strings包
	fmt.Println(strings.Contains(a, "h"))
	fmt.Println(strings.Count(a, "h"))
	fmt.Println(strings.Split(a, "e"))
	fmt.Println(strings.HasPrefix(a, "h"))
	fmt.Println(strings.Index(a, "o"))

}
