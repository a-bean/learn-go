package main

import "fmt"

// Duck interface定义
type Duck interface {
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (p *pskDuck) Gaga() {
	fmt.Println("gaga")
}
func (p *pskDuck) Walk() {
	fmt.Println("Walk")
}
func (p *pskDuck) Swimming() {
	fmt.Println("Swimming")
}

// MyWriter 多接口
type MyWriter interface {
	Write() error
}
type MyCloser interface {
	Close() error
}

type writerCloser struct {
}

func (p *writerCloser) Write() error {
	fmt.Println("Write")
	return nil
}
func (p *writerCloser) Close() error {
	fmt.Println("Close")
	return nil
}

/* 接口类型断言 */
func add(a, b any) any {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case int32:
		return a.(int32) + b.(int32)
	case int64:
		return a.(int64) + b.(int64)
	case string:
		return a.(string) + b.(string)
	default:
		panic("not support type")
	}
}

// MyBoth 接口嵌套
type MyBoth interface {
	MyWriter
	MyCloser
}

func main() {
	// go语言中处处都是鸭子类型(强调的是外部行为：方法)

	var d Duck = &pskDuck{}
	d.Walk()

	// 多接口
	var myw MyWriter = &writerCloser{}
	var myc MyCloser = &writerCloser{}
	var myb MyBoth = &writerCloser{}
	fmt.Println(myw, myc, myb)

	/* 接口类型断言 */

}
