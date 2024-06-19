package queue

import (
	"container/list"
	"fmt"
)

type LQueue struct {
	queue *list.List
}

func (lq *LQueue) Enqueue(value any) {
	lq.queue.PushBack(value)
}

func (lq *LQueue) Dequeue() error {

	if !lq.Empty() {
		element := lq.queue.Front()
		lq.queue.Remove(element)

		return nil
	}

	return fmt.Errorf("dequeue is empty we got an error")
}

func (lq *LQueue) Front() (any, error) {
	if !lq.Empty() {
		val := lq.queue.Front().Value
		return val, nil
	}

	return "", fmt.Errorf("error queue is empty")
}

func (lq *LQueue) Back() (any, error) {
	if !lq.Empty() {
		val := lq.queue.Back().Value
		return val, nil
	}

	return "", fmt.Errorf("error queue is empty")
}

func (lq *LQueue) Len() int {
	return lq.queue.Len()
}

func (lq *LQueue) Empty() bool {
	return lq.queue.Len() == 0
}
