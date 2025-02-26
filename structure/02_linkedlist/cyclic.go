package linkedlist

import "fmt"

type Cyclic[T any] struct {
	Size int
	Head *Node[T]
}

func NewCyclic[T any]() *Cyclic[T] {
	return &Cyclic[T]{}
}

func (cl *Cyclic[T]) Add(val T) {
	n := NewNode(val)
	cl.Size++
	if cl.Head == nil {
		n.Prev = n
		n.Next = n
		cl.Head = n
	} else {
		n.Prev = cl.Head.Prev
		n.Next = cl.Head
		cl.Head.Prev.Next = n
		cl.Head.Prev = n
	}
}

func (cl *Cyclic[T]) Rotate(places int) {
	if cl.Size > 0 {
		if places < 0 {
			multiple := cl.Size - 1 - places/cl.Size
			places += multiple * cl.Size
		}
		places %= cl.Size

		if places > cl.Size/2 {
			places = cl.Size - places
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Prev
			}
		} else if places == 0 {
			return
		} else {
			for i := 0; i < places; i++ {
				cl.Head = cl.Head.Next
			}

		}
	}
}

func (cl *Cyclic[T]) Delete() bool {
	var deleted bool
	var prevItem, thisItem, nextItem *Node[T]

	if cl.Size == 0 {
		return deleted
	}

	deleted = true
	thisItem = cl.Head
	nextItem = thisItem.Next
	prevItem = thisItem.Prev

	if cl.Size == 1 {
		cl.Head = nil
	} else {
		cl.Head = nextItem
		nextItem.Prev = prevItem
		prevItem.Next = nextItem
	}
	cl.Size--

	return deleted
}

func (cl *Cyclic[T]) Destroy() {
	for cl.Delete() {
		continue
	}
}

func (cl *Cyclic[T]) Walk() *Node[T] {
	var start *Node[T]
	start = cl.Head

	for i := 0; i < cl.Size; i++ {
		fmt.Printf("%v \n", start.Val)
		start = start.Next
	}
	return start
}

func JosephusProblem(cl *Cyclic[int], k int) int {
	for cl.Size > 1 {
		cl.Rotate(k)
		cl.Delete()
		cl.Rotate(-1)
	}
	retval := cl.Head.Val
	cl.Destroy()
	return retval
}
