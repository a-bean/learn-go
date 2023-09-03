package linkedlist

import "fmt"

type Doubly[T any] struct {
	Head *Node[T]
}

func (ll *Doubly[T]) Init() *Doubly[T] {
	ll.Head = &Node[T]{}
	ll.Head.Next = ll.Head
	ll.Head.Prev = ll.Head

	return nil
}

func NewDoubly[T any]() *Doubly[T] {
	return new(Doubly[T]).Init()
}

func (ll *Doubly[T]) lazyInit() {
	if ll.Head.Next == nil {
		ll.Init()
	}
}

// Count 数量
func (ll *Doubly[T]) Count() int {
	if ll.Head.Next == nil {
		return 0
	}

	ctr := 0
	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		ctr += 1
	}
	return ctr
}

// insert
func (ll *Doubly[T]) insert(n, at *Node[T]) *Node[T] {
	n.Prev = at
	n.Next = at.Prev
	n.Prev.Next = n
	n.Next.Prev = n

	return n
}

func (ll *Doubly[T]) insertValue(val T, at *Node[T]) *Node[T] {
	return ll.insert(NewNode(val), at)
}

func (ll *Doubly[T]) AddAtBeg(val T) {
	ll.lazyInit()
	ll.insertValue(val, ll.Head)
}

func (ll *Doubly[T]) AddAtEnd(val T) {
	ll.lazyInit()
	ll.insertValue(val, ll.Head.Prev)
}

// Remove 移除元素
func (ll *Doubly[T]) Remove(n *Node[T]) T {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil
	n.Prev = nil
	return n.Val
}

func (ll *Doubly[T]) DelAtBeg() (T, bool) {
	if ll.Head.Next == nil {
		var r T
		return r, false
	}

	n := ll.Head.Next
	val := n.Val
	ll.Remove(n)
	return val, true
}

func (ll *Doubly[T]) DelAtEnd() (T, bool) {
	if ll.Head.Next == nil {
		var r T
		return r, false
	}

	n := ll.Head.Prev
	val := n.Val
	ll.Remove(n)
	return val, true
}

func (ll *Doubly[T]) DelByPos(pos int) (T, bool) {
	switch {
	case ll.Head == nil:
		var r T
		return r, false
	case pos-1 == 0:
		return ll.DelAtBeg()
	case pos-1 == ll.Count():
		return ll.DelAtEnd()
	case pos-1 > ll.Count():
		var r T
		return r, false
	}
	var prev *Node[T]
	var val T
	cur := ll.Head
	count := 0
	for count < pos-1 {
		prev = cur
		cur = cur.Next
		count++
	}
	cur.Next.Prev = prev
	val = cur.Val
	prev.Next = cur.Next

	return val, true

}

func (ll *Doubly[T]) Reverse() {
	var Prev, Next *Node[T]
	cur := ll.Head

	for cur != nil {
		Next = cur.Next
		cur.Next = Prev
		cur.Prev = Next
		Prev = cur
		cur = Next
	}
	ll.Head = Prev
}

func (ll *Doubly[T]) Display() {
	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		fmt.Print(cur.Val, " ")
	}
	fmt.Print("\n")

}
