package main

import "fmt"

// 代理模式为其他对象提供一种代理，以控制对这个对象的访问。

// Image 接口
type Image interface {
	Display()
}

// RealImage 具体对象
type RealImage struct {
	filename string
}

func (r *RealImage) Display() {
	fmt.Println("Displaying", r.filename)
}

func NewRealImage(filename string) *RealImage {
	fmt.Println("Loading", filename)
	return &RealImage{filename: filename}
}

// ProxyImage 代理
type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

func main() {
	image := &ProxyImage{filename: "test.jpg"}

	// 图像将加载并显示
	image.Display()

	// 图像不会再次加载，直接显示
	image.Display()
}
