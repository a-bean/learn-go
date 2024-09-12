package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// 长度
	name := "kobe"
	bytes := []byte(name)
	byter := []rune(name)
	char1 := '1'
	char3 := '哦'
	var char4 byte = '1'
	var char5 rune = '1'
	fmt.Println(char1, char3, char4, char5)
	fmt.Println(reflect.TypeOf(char1), reflect.TypeOf(char3), reflect.TypeOf(char4), reflect.TypeOf(char5))
	fmt.Println(len(name), bytes, byter)
	fmt.Println(len(name), len(bytes), len(byter))

	// rune字面量是指通过\,\x,\u,\U开头的字符串,分别对应8进制,16进制,unicode.注意\u和\U开头的字符串都是
	//通过unicode码表示的.只是\U开头的字符串表示的码范围更宽
	rune1 := "\x61\x62\x63"
	rune2 := "\x61"
	rune3 := "\u0061\u0062\u0063"
	rune4 := "\U00000061\U00000062\U00000063"

	fmt.Println(rune1, reflect.TypeOf(rune1), rune2)
	fmt.Println(rune3, reflect.TypeOf(rune1), rune4)
	// 双引号跟反引号本质区别是双引号的字符串支持转义,而反应好的字符串不支持
	// 转义符
	curseName := "faf\"efa\""
	curseName1 := `faf"efa"`
	fmt.Println(curseName)
	fmt.Println(curseName1)

	// 字符串比较
	a := "hello"
	b := "good"
	fmt.Println(a == b)
	fmt.Println(a > b) // 比较ASCII

	// strings包
	fmt.Println(strings.Contains(a, "h"))
	fmt.Println(strings.Count(a, "h"))
	fmt.Println(strings.Split(a, "e"))
	fmt.Println(strings.HasPrefix(a, "h"))
	fmt.Println(strings.Index(a, "o"))

	// Trim,TrimLeft,TrimRight将检查每个开头或者结尾的rune值,直到遇到一个不满足条件的rune值为止
	// TrimPrefix,TrimSuffix是只遇到一次满足子串条件的rune值就会终止
	var s = "go123gogogogo"
	fmt.Println(strings.TrimRight(s, "go"))
	fmt.Println(strings.Trim(s, "go"))
	fmt.Println(strings.TrimLeft(s, "go"))

	fmt.Println(strings.TrimSuffix(s, "go"))
	fmt.Println(strings.TrimPrefix(s, "go"))

	var aaa uint64 = math.MaxUint64
	var bbb uint64 = 1
	fmt.Println("aa", aaa+bbb)
	//fmt.Println(math.MaxInt32 + 2)
	//fmt.Println(math.MaxInt32 + math.MaxInt32)
}
