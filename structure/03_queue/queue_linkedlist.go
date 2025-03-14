package queue

type Node struct {
	Data any
	Next *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *Queue) enqueue(n any) {
	var newNode Node // create new Node
	newNode.Data = n // set the data

	if ll.tail != nil {
		ll.tail.Next = &newNode
	}

	ll.tail = &newNode

	if ll.head == nil {
		ll.head = &newNode
	}
	ll.length++
}

func (ll *Queue) dequeue() any {
	if ll.isEmpty() {
		return -1 // if is empty return -1
	}
	data := ll.head.Data

	ll.head = ll.head.Next

	if ll.head == nil {
		ll.tail = nil
	}

	ll.length--
	return data
}

func (ll *Queue) isEmpty() bool {
	return ll.length == 0
}

func (ll *Queue) len() int {
	return ll.length
}

func (ll *Queue) frontQueue() any {
	return ll.head.Data
}

func (ll *Queue) backQueue() any {
	return ll.tail.Data
}
