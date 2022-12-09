package linkedlist

type Singly[T any] struct {
	length int
	Head   *Node[T]
}

func NewSingly[T any]() *Singly[T] {
	return &Singly[T]{}
}

func (ll *Singly[T]) AddAtBeg(val T) {
	n := NewNode(val)
	n.Next = ll.Head
	ll.Head = n
	ll.length++
}

func (ll *Singly[T]) AddAtEnd(val T) {
	n := NewNode(val)

	if ll.Head == nil {
		ll.Head = n
		ll.length++
		return
	}

	cur := ll.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
	ll.length++
}

func (ll *Singly[T]) DelAtBeg() (T, bool) {
	if ll.Head == nil {
		var r T
		return r, false
	}

	cur := ll.Head
	ll.Head = cur.Next
	ll.length--
	return cur.Val, true
}

func (ll *Singly[T]) DelAtEnd() (T, bool) {
	if ll.Head == nil {
		var r T
		return r, false
	}

	if ll.Head.Next == nil {
		return ll.DelAtBeg()
	}

	cur := ll.Head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}

	retVal := cur.Next.Val
	cur.Next = nil
	ll.length--

	return retVal, true

}
