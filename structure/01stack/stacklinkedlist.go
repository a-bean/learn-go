package stack

import "fmt"

type Node struct {
	Val  any
	Next *Node
}

type ListStack struct {
	top    *Node
	length int
}

func (ls *ListStack) Push(n any) {
	node := &Node{
		Val:  n,
		Next: ls.top,
	}
	ls.top = node
	ls.length++
}

func (ls *ListStack) Pop() any {
	cur := ls.top.Val
	if ls.top.Next == nil {
		ls.top = nil
	} else {
		ls.top.Val, ls.top.Next = ls.top.Next.Val, ls.top.Next.Next
	}
	ls.length--
	return cur
}

func (ls *ListStack) Length() int {
	return ls.length
}

func (ls *ListStack) IsEmpty() bool {
	return ls.length == 0
}

func (ls *ListStack) Peak() any {
	return ls.top.Val
}

func (ls *ListStack) Show() []any {
	var list []any
	cur := ls.top
	if cur.Val != nil {
		list = append(list, cur.Val)
		cur = cur.Next
		fmt.Println(cur)
	}
	return list
}
