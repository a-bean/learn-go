package linkedlist

import "fmt"

type Doubly[T any] struct {
	Head *Node[T]
}

func (ll *Doubly[T]) Init() *Doubly[T] {
	ll.Head = &Node[T]{}
	ll.Head.Next = ll.Head
	ll.Head.Prev = ll.Head

	return ll
}

func NewDoubly[T any]() *Doubly[T] {
	return new(Doubly[T]).Init()
}

func (ll *Doubly[T]) lazyInit() {
	if ll.Head.Next == nil {
		ll.Init()
	}
}

func (ll *Doubly[T]) insert(n, at *Node[T]) *Node[T] {
	n.Prev = at
	n.Next = at.Next
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

func (ll *Doubly[T]) Remove(n *Node[T]) T {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil
	n.Prev = nil

	return n.Val
}

func (ll *Doubly[T]) DelAtBeg() (T, bool) {
	// no item
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
	// no item
	if ll.Head.Prev == nil {
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

func (ll *Doubly[T]) Count() int {
	var ctr int = 0

	if ll.Head.Next == nil {
		return 0
	}

	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		ctr += 1
	}

	return ctr
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

func (ll *Doubly[T]) DisplayReverse() {
	if ll.Head == nil {
		return
	}
	var cur *Node[T]
	for cur = ll.Head.Prev; cur != ll.Head; cur = cur.Prev {
		fmt.Print(cur.Val, " ")
	}

	fmt.Print("\n")
}

func (ll *Doubly[T]) Front() *Node[T] {
	if ll.Count() == 0 {
		return nil
	}

	return ll.Head.Next
}

func (ll *Doubly[T]) Back() *Node[T] {
	if ll.Count() == 0 {
		return nil
	}

	return ll.Head.Prev
}

func (ll *Doubly[T]) MoveToBack(n *Node[T]) {
	if ll.Head.Prev == n {
		return
	}

	ll.move(n, ll.Head.Prev)
}

func (ll *Doubly[T]) move(n, at *Node[T]) {
	if n == at {
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Prev = at
	n.Next = at.Next
	n.Prev.Next = n
	n.Next.Prev = n
}
