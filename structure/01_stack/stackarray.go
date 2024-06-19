package stack

var stackArray []any

func Push(n any) {
	stackArray = append([]any{n}, stackArray...)
}

func Length() int {
	return len(stackArray)
}

func Peak() any {
	return stackArray[0]
}

func Pop() any {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}

func IsEmpty() bool {
	return len(stackArray) == 0
}

// Stack struct形式
type Stack struct {
	container []any
}

func (s *Stack) Push(n any) {
	s.container = append([]any{n}, s.container...)
}

func (s *Stack) Length() int {
	return len(s.container)
}

func (s *Stack) Peak() any {
	return s.container[0]
}

func (s *Stack) Pop() any {
	pop := s.container[0]
	s.container = s.container[1:]
	return pop
}

func (s *Stack) IsEmpty() bool {
	return len(s.container) == 0
}
