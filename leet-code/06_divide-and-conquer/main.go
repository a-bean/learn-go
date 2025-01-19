package main

import "fmt"

// 22. 括号生成 https://leetcode.cn/problems/generate-parentheses/description/

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

// 50 Pow(x, n) https://leetcode.cn/problems/powx-n/description/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

func main() {
	fmt.Println(generateParenthesis(1))
}
