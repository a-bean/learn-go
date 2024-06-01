package main

import "fmt"

// https://leetcode.cn/problems/generate-parentheses/description/

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	store := make(map[int][]string)
	if ans, ok := store[n]; ok {
		return ans
	}
	ans := []string{}
	for k := 1; k <= n; k++ {
		A := generateParenthesis(k - 1)
		B := generateParenthesis(n - k)
		for _, a := range A {
			for _, b := range B {
				ans = append(ans, "("+a+")"+b)
			}
		}
	}
	store[n] = ans
	return ans
}

func main() {
	fmt.Println(generateParenthesis(1))
}
