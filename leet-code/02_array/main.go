package main

import (
	"fmt"
)

// 26. 删除有序数组中的重复项 https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

// 283. 移动零 https://leetcode.cn/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}

	for j := index; j < len(nums); j++ {
		nums[j] = 0
	}
}

// 88. 合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/description/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n; m > 0 && n > 0; i-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[i-1] = nums1[m-1]
			m--
		} else {
			nums1[i-1] = nums2[n-1]
			n--
		}
	}

	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}

// 1. 两数之和 https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 53. 最大子数组和 https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := 0
	for i := 0; i < len(nums); i++ {
		if curSum < 0 {
			curSum = 0
		}
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

// 73. 矩阵置零 https://leetcode.cn/problems/set-matrix-zeroes/description/
func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

// 54. 螺旋矩阵 https://leetcode.cn/problems/spiral-matrix/description/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := []int{}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}

		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}

		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}

// 48. 旋转图像 https://leetcode.cn/problems/rotate-image/description/
func rotate(matrix [][]int) {
	n := len(matrix)
	// 先沿着对角线翻转
	//[[1, 2, 3],				[[1, 4, 7],
	// [4, 5, 6],  ==>   [2, 5, 8],
	// [7, 8, 9]]        [3, 6, 9]]
	//
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 再沿着中线翻转
	//[[1, 4, 7],				[[7, 4, 1],
	// [2, 5, 8],  ==>   [8, 5, 2],
	// [3, 6, 9]]        [9, 6, 3]]
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// 240. 搜索二维矩阵 II https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
// 思路 是从右上角开始搜索，当前元素大于目标值则向左移动，当前元素小于目标值则向下移动
// 直到找到目标值或者越界
// 复杂度分析：时间复杂度 O(m+n) 空间复杂度 O(1)
func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	x, y := 0, m-1
	for x < n && y >= 0 {
		if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		} else {
			return true
		}
	}
	return false
}

// 暴力破解法
// 复杂度分析：时间复杂度 O(m*n) 空间复杂度 O(1)
func searchMatrix1(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 560 // 和为 K 的子数组 https://leetcode.cn/problems/subarray-sum-equals-k/description/
func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	m := make(map[int]int)
	m[0] = 1 // 初始化前缀和为0的情况
	for _, num := range nums {
		sum += num
		if v, ok := m[sum-k]; ok {
			count += v // 如果存在前缀和为 sum-k 的情况，则说明找到了一个子数组
		}
		m[sum]++ // 更新前缀和的计数
	}
	return count
}

// 238 // 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/description/
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// answer[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// R 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	R := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		answer[i] = answer[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= nums[i]
	}
	return answer
}

// 41 // 缺失的第一个正数 https://leetcode.cn/problems/first-missing-positive/description/
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 将 nums[i] 放到正确的位置上 即 nums[i] 应该在索引 nums[i]-1 的位置上
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	// 遍历数组，找到第一个不满足条件的索引
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1 // 如果所有位置都满足条件，则返回 n+1
}

func main() {
	fmt.Print(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	moveZeroes([]int{1, 1, 1, 0, 2, 2, 3})
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	fmt.Print(twoSum([]int{1, 2, 3, 4}, 5))
	fmt.Print(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	productExceptSelf([]int{1, 2, 3, 4})
	firstMissingPositive([]int{1, 2, 0})
	firstMissingPositive([]int{-1, 4, 2, 1, 9, 10})
	spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}
