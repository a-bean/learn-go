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

// https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/ 返回 A 的最短的非空连续子数组的长度，该子数组的和 至少 为 K 。如果没有和至少为 K 的非空子数组，返回 -1
func shortestSubarray(A []int, K int) int {
	res, prefixSum := len(A)+1, make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		prefixSum[i+1] = prefixSum[i] + A[i]
	}
	// deque 中保存递增的 prefixSum 下标
	deque := []int{}
	for i := range prefixSum {
		// 下面这个循环希望能找到 [deque[0], i] 区间内累加和 >= K，如果找到了就更新答案
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= K {
			length := i - deque[0]
			if res > length {
				res = length
			}
			// 找到第一个 deque[0] 能满足条件以后，就移除它，因为它是最短长度的子序列了
			deque = deque[1:]
		}
		// 下面这个循环希望能保证 prefixSum[deque[i]] 递增
		for len(deque) > 0 && prefixSum[i] <= prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if res <= len(A) {
		return res
	}
	return -1
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(shortestSubarray([]int{1, 3, -1, -3, 5, 3, 6, 7}, 2))

}
