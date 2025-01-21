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

// 23 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	min := length / 2
	left := mergeKLists(lists[:min])
	right := mergeKLists(lists[min:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Val < right.Val {
		left.Next = mergeTwoLists(left.Next, right)
		return left
	}
	right.Next = mergeTwoLists(left, right.Next)
	return right

}

func main() {
	fmt.Println(generateParenthesis(1))
	mergeKLists([]*ListNode{{Val: 1, Next: nil}, {Val: 2, Next: nil}, {Val: 3, Next: nil}})
}
