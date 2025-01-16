package main

import (
	"fmt"
)

// 前缀和 1248 53 304

// 1248 优美子数组 https://leetcode.cn/problems/count-number-of-nice-subarrays/description/
func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n+1) // 前缀和数组
	for i := 1; i <= n; i++ {
		// 统计奇数个数前缀和
		sum[i] = sum[i-1] + nums[i-1]%2
	}

	ans := 0
	count := make([]int, n+1) // 统计每个前缀和出现的次数
	count[0]++                // 初始化前缀和为0的出现次数

	// 遍历，使用双指针法
	for i := 1; i <= n; i++ {
		// 当前前缀和
		count[sum[i]]++

		// 计算以当前前缀和为基础的 k 个奇数的子数组
		if sum[i] >= k {
			ans += count[sum[i]-k] // 如果存在前缀和sum[i]-k，则增加结果
		}
	}

	return ans // 返回结果
}

// 53 最大子序和 https://leetcode.cn/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0 // 如果数组为空，则返回0
	}

	// 创建前缀和数组
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]

	// 计算前缀和
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	maxSum := prefixSum[0] // 初始化最大和为第一个前缀和

	// 遍历前缀和数组以找到最大子数组和
	for end := 0; end < n; end++ {
		for start := 0; start <= end; start++ {
			var currentSum int
			if start == 0 {
				currentSum = prefixSum[end] // 从头开始
			} else {
				currentSum = prefixSum[end] - prefixSum[start-1] // 使用前缀和计算当前子数组和
			}
			maxSum = max(maxSum, currentSum) // 更新最大和
		}
	}

	return maxSum // 返回结果
}

// 辅助函数，用于返回两个整数中的较大者
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 304 二维区域和检索 - 矩阵不可变 https://leetcode.cn/problems/range-sum-query-2d-immutable/description/
// 二维数组前缀和
type NumMatrix struct {
	cumsum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{nil}
	}
	cumsum := make([][]int, len(matrix)+1)
	cumsum[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		cumsum[i+1] = make([]int, len(matrix[i])+1)
		for j := range matrix[i] {
			cumsum[i+1][j+1] = matrix[i][j] + cumsum[i][j+1] + cumsum[i+1][j] - cumsum[i][j]
		}
	}
	return NumMatrix{cumsum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	cumsum := this.cumsum
	return cumsum[row2+1][col2+1] - cumsum[row1][col2+1] - cumsum[row2+1][col1] + cumsum[row1][col1]
}

// 差分 1109

// 1109 航班预订统计 https://leetcode.cn/problems/corporate-flight-bookings/description/
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))

}
