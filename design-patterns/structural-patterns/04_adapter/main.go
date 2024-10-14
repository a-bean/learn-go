// 适配器模式将一个类的接口转换成客户希望的另一个接口。
package main

import "fmt"

// 目标接口
type Target interface {
	Request() string
}

// Adaptee 适配者
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's Specific Request"
}

// Adapter 适配器
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{adaptee: adaptee}

	fmt.Println(adapter.Request())
}
