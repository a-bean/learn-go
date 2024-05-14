package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)

	for _, value := range s {
		if value == '(' || value == '[' || value == '{' {

			stack = append(stack, value)

		} else if len(stack) > 0 && value == ')' && stack[len(stack)-1] == '(' || len(stack) > 0 && value == ']' && stack[len(stack)-1] == '[' || len(stack) > 0 && value == '}' && stack[len(stack)-1] == '{' {

			stack = stack[:len(stack)-1]

		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	println(isValid("{}()[]"))
}
