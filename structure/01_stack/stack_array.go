package stack

var stackArray []any

func stackPush(n any) {
	stackArray = append([]any{n}, stackArray...)
}

func stackLength() int {
	return len(stackArray)
}

func stackPeak() any {
	return stackArray[0]
}

func stackEmpty() bool {
	return len(stackArray) == 0
}

func stackPop() any {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}
