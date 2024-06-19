package stack

import "fmt"

type Node struct {
	Val  any
	Next *Node
}

type ListStack struct {
	head   *Node
	length int
}

func (ls *ListStack) Push(n any) {
	node := &Node{
		Val:  n,
		Next: ls.head,
	}
	ls.head = node
	ls.length++
}

func (ls *ListStack) Pop() any {
	cur := ls.head.Val
	if ls.head.Next == nil {
		ls.head = nil
	} else {
		ls.head.Val, ls.head.Next = ls.head.Next.Val, ls.head.Next.Next
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
	return ls.head.Val
}

func (ls *ListStack) Show() []any {
	var list []any
	cur := ls.head
	if cur.Val != nil {
		list = append(list, cur.Val)
		cur = cur.Next
		fmt.Println(cur)
	}
	return list
}
