package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/valid-parentheses/
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

// https://leetcode.cn/problems/simplify-path/
func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	stack := make([]string, 0)
	var res string
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			stack = append(stack, cur)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return res
}

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, v)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}

	return stack[0]
}

func main() {
	fmt.Println(isValid("{}()[]"))
	fmt.Println(simplifyPath("///d//da///da"))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
