package main

import "fmt"

// https://leetcode.cn/problems/sliding-window-maximum/description/
func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0, k)
	ans := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 删除出界的选项
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// 删除比当前值小的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		// 插入新选项并维护单调性
		deque = append(deque, i)

		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

}
