package stack

import (
	"container/list"
	"fmt"
)

type SList struct {
	stack *list.List
}

func (sl *SList) Push(n any) {
	sl.stack.PushFront(n)
}

func (sl *SList) Pop() (any, error) {
	if sl.IsEmpty() {
		return "", fmt.Errorf("this stack is empty")
	}
	element := sl.stack.Front()
	sl.stack.Remove(element)
	return element.Value, nil
}

func (sl *SList) Length() int {
	return sl.stack.Len()
}

func (sl *SList) IsEmpty() bool {
	return sl.stack.Len() == 0
}

func (sl *SList) Peak() (any, error) {
	if sl.IsEmpty() {
		return "", fmt.Errorf("this stack is empty")
	}
	element := sl.stack.Front()
	return element.Value, nil
}
