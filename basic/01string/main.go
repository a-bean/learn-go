package main

import (
	bytes2 "bytes"
	"fmt"
	"math"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
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
	// 双引号跟反引号本质区别是双引号的字符串支持转义,而反引号的字符串不支持
	// 转义符
	curseName := "faf\"efa\""
	curseName1 := `faf\"efa\"`
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

	// 字符串的len(): 是字符串的字节长度,而非字符个数
	s1 := "abc"
	s2 := "中文"
	fmt.Println(len(s1), unsafe.Sizeof(s1), utf8.RuneCountInString(s1), len([]rune(s1)))
	fmt.Println(len(s2), unsafe.Sizeof(s2), utf8.RuneCountInString(s2), len([]rune(s2)))

	// 字符串的拼接方式
	s3 := s1 + s2 // 会创建一个新的空间 旧的对象进行gc
	s4 := fmt.Sprintf("%s%s", s1, s2)

	var buf bytes2.Buffer
	buf.Grow(len(s1) + len(s2))
	buf.WriteString(s1)
	buf.WriteString(s2)

	// 推荐
	// 性能略高于bytes2.Buffer,bytes2.Buffer转换字符串需要重新申请内存空间,
	// strings.Builder是将底层的bytes转换为string
	var bul strings.Builder
	bul.Grow(len(s1) + len(s2))
	bul.WriteString(s1)
	bul.WriteString(s2)

	bt := make([]byte, 0, len(s1)+len(s2))
	bt = append(bt, s1...)
	bt = append(bt, s2...)

	sl := []string{s1, s2}
	st := strings.Join(sl, "")

	fmt.Println(s3, s4, buf.String(), bul.String(), string(bt), st)

	var ss1 = "kebo"
	ss2 := ss1
	ss3 := ss1
	fmt.Println(&ss1, &ss2, &ss3)
	prt := (*reflect.StringHeader)(unsafe.Pointer(&ss1))
	prt2 := (*reflect.StringHeader)(unsafe.Pointer(&ss2))            // 字符串在赋值的过程中不会发生拷贝
	fmt.Println(unsafe.Pointer(prt.Data), unsafe.Pointer(prt2.Data)) // 两个相同(赋值底层指针不变)

}

/*
	避免字符串截取而导致内存泄漏
	1. 将字符串转成字节切片,再转成字符串
	2. 截取后在前面拼接一个新的字符串,然后在截掉
	3. 使用strings.Builder对新的字符串进行重构

*/

func sliceString(s string, index int) string {
	return string([]byte(s[:index])) // 这种强转的方式会进行拷贝(开辟新的空间),跟原始值没有引用关系
}
