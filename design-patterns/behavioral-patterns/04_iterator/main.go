package main

import "fmt"

// 迭代器模式: 提供一种方法顺序访问集合中的各个元素，而不暴露其底层表示

// Iterator 集合接口
type Iterator interface {
	HasNext() bool
	Next() string
}

// NameIterator 具体集合
type NameIterator struct {
	names []string
	index int
}

func (n *NameIterator) HasNext() bool {
	return n.index < len(n.names)
}

func (n *NameIterator) Next() string {
	if n.HasNext() {
		name := n.names[n.index]
		n.index++
		return name
	}
	return ""
}

func main() {
	names := []string{"John", "Jane", "Jack", "Jill"}
	iterator := &NameIterator{names: names}

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
