package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 20: 有效的括号 https://leetcode.cn/problems/valid-parentheses/
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

// 71: 简化路径 https://leetcode.cn/problems/simplify-path/
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

// 150: 逆波兰表达式求值（后缀表达式） https://leetcode.cn/problems/evaluate-reverse-polish-notation/
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

// 227: 基本计算器 II https://leetcode.cn/problems/basic-calculator-ii/
func calculateII(s string) int {
	tokens := make([]string, 0)
	ops := make([]byte, 0)
	num := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num += string(s[i])
		} else {
			// 处理累积的数字
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}

			if s[i] == ' ' {
				continue
			}

			// 只处理运算符
			if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
				currentRank := getRank(s[i])
				for len(ops) > 0 && getRank(ops[len(ops)-1]) >= currentRank {
					tokens = append(tokens, string(ops[len(ops)-1]))
					ops = ops[:len(ops)-1]
				}
				ops = append(ops, s[i])
			}
		}
	}

	// 处理最后一个数字
	if num != "" {
		tokens = append(tokens, num)
	}

	// 处理剩余的运算符
	for len(ops) > 0 {
		tokens = append(tokens, string(ops[len(ops)-1]))
		ops = ops[:len(ops)-1]
	}

	return evalRPN(tokens)
}

func getRank(s byte) int {
	if s == '*' || s == '/' {
		return 2
	}
	if s == '+' || s == '-' {
		return 1
	}
	return 0
}

// 224: 基本计算器 https://leetcode.cn/problems/basic-calculator/
// func calculate(s string) int {

// }

// 84: 柱状图中最大的矩形 https://leetcode.cn/problems/largest-rectangle-in-histogram/description/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type rect struct {
	width  int
	height int
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 保证最后stack全部弹出
	stack := make([]rect, len(heights)/2)
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		width := 0

		for len(stack) > 0 && stack[len(stack)-1].height >= heights[i] {
			width += stack[len(stack)-1].width
			maxArea = max(maxArea, stack[len(stack)-1].height*width)
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, rect{width: width + 1, height: heights[i]})
	}

	return maxArea
}

// 42: 接雨水 https://leetcode.cn/problems/trapping-rain-water/
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func trap(height []int) int {
	area := 0
	stack := make([]rect, 0)
	for i := 0; i < len(height); i++ {
		width := 0
		for len(stack) > 0 && stack[len(stack)-1].height <= height[i] {
			bottom := stack[len(stack)-1].height
			width += stack[len(stack)-1].width
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				continue
			}
			top := min(stack[len(stack)-1].height, height[i])
			area += width * (top - bottom)
		}
		stack = append(stack, rect{width: width + 1, height: height[i]})
	}
	return area
}

// 496: 下一个更大元素 I https://leetcode.cn/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	res, record := []int{}, map[int]int{}
	for k, v := range nums2 {
		record[v] = k
	}

	for i := 0; i < len(nums1); i++ {
		flag := false
		for j := record[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				res = append(res, nums2[j])
				flag = true
				break
			}
		}
		if !flag {
			res = append(res, -1)
		}
	}

	return res
}

func main() {
	fmt.Println(isValid("{}()[]"))
	fmt.Println(simplifyPath("///d//da///da"))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Println(largestRectangleArea([]int{1, 2, 3, 4, 5}))
	fmt.Println(trap([]int{2, 1, 2}))
	nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
}
