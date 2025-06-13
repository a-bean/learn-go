package main

import (
	"fmt"
	"sort"
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

// 11 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water/description/
// 双指针
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max

}

// 递归 78 77 46 47

// 78 子集  https://leetcode.cn/problems/subsets/description/
// subsets 返回给定整数数组的所有子集
// 输入：nums = [1,2,3]
// 输出：[[],[1],[1,2],[1,2,3],[1,3],[2],[2,3],[3]]
func subsets(nums []int) [][]int {
	result := [][]int{} // 存放所有子集的结果
	path := []int{}     // 当前子集路径
	var backtrack func(start int)

	backtrack = func(start int) {
		// 将当前子集添加到结果中
		result = append(result, append([]int{}, path...))

		// 遍历数组元素
		for i := start; i < len(nums); i++ {
			// 选择当前元素
			path = append(path, nums[i])
			// 递归调用
			backtrack(i + 1)
			// 撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrack(0) // 从第一个元素开始生成所有子集
	return result
}

//	17// 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
//
// letterCombinations 返回给定电话号码数字的所有字母组合

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{} // 如果输入为空，返回空切片
	}

	var letterMap = []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	result := []string{} // 存放所有组合的结果
	path := []byte{}     // 当前组合路径

	var backtrack func(index int)

	backtrack = func(index int) {
		// 当路径长度等于 digits 的长度时，添加到结果中
		if index == len(digits) {
			result = append(result, string(path)) // 复制当前路径
			return
		}

		// 获取当前数字对应的字母
		num := digits[index] - '0'
		letters := letterMap[num] // 获取对应的字母字符串

		// 遍历字母
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i]) // 选择当前字母
			backtrack(index + 1)            // 递归调用，处理下一个数字
			path = path[:len(path)-1]       // 撤销选择
		}
	}

	backtrack(0) // 从第一个数字开始生成组合
	return result
}

// 77 组合  https://leetcode.cn/problems/combinations/description/
// combine 返回从 1 到 n 中选择 k 个数的所有组合
// 输入：n = 4, k = 2
// 输出：[[2,4],[3,4],[2,3],[1,2],[1,3],[1,4]]
func combine(n int, k int) [][]int {
	result := [][]int{} // 存放所有组合的结果
	path := []int{}     // 当前组合路径
	var backtrack func(start int)

	backtrack = func(start int) {
		// 当当前组合长度达到 k 时，保存组合
		if len(path) == k {
			result = append(result, append([]int{}, path...))
			return
		}

		// 遍历选择
		for i := start; i <= n; i++ {
			// 选择当前数
			path = append(path, i)
			// 递归选择下一个数
			backtrack(i + 1)
			// 撤销选择
			path = path[:len(path)-1]
		}
	}

	backtrack(1) // 从1开始生成组合
	return result
}

// 39 组合总和 https://leetcode.cn/problems/combination-sum/description/
// combinationSum 返回给定整数数组中所有和为 target 的组合
func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{} // 存放所有组合的结果
	path := []int{}     // 当前组合路径

	var backtrack func(start int, remaining int)

	backtrack = func(start int, remaining int) {
		// 如果剩余值为0，说明找到一个组合
		if remaining == 0 {
			result = append(result, append([]int{}, path...)) // 复制当前路径
			return
		}

		// 如果剩余值小于0，说明当前路径不合法
		if remaining < 0 {
			return
		}

		// 遍历候选数字
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])    // 选择当前数字
			backtrack(i, remaining-candidates[i]) // 递归调用，允许重复使用同一数字
			path = path[:len(path)-1]             // 撤销选择
		}
	}

	backtrack(0, target) // 从第一个数字开始生成组合
	return result
}

// 22 括号生成 https://leetcode.cn/problems/generate-parentheses/description/
// generateParenthesis 返回所有可能的括号组合
func generateParenthesis(n int) []string {
	result := []string{} // 存放所有括号组合的结果
	path := []byte{}     // 当前括号组合路径

	var backtrack func(left int, right int)

	backtrack = func(left int, right int) {
		// 如果左括号和右括号都用完了，说明找到一个合法组合
		if left == 0 && right == 0 {
			result = append(result, string(path)) // 复制当前路径
			return
		}

		// 如果左括号还有剩余，可以添加左括号
		if left > 0 {
			path = append(path, '(')  // 选择左括号
			backtrack(left-1, right)  // 递归调用
			path = path[:len(path)-1] // 撤销选择
		}

		// 如果右括号的数量大于左括号，可以添加右括号
		if right > left {
			path = append(path, ')')  // 选择右括号
			backtrack(left, right-1)  // 递归调用
			path = path[:len(path)-1] // 撤销选择
		}
	}

	backtrack(n, n) // 从 n 个左括号和 n 个右括号开始生成组合
	return result
}

// 79 单词搜索 https://leetcode.cn/problems/word-search/description/
// exist 检查给定的单词是否存在于二维网格中
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	directions := []struct{ dx, dy int }{
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
	}

	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		if index == len(word) {
			return true // 找到完整单词
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[index] {
			return false // 越界或字符不匹配
		}

		temp := board[x][y] // 保存当前字符
		board[x][y] = '#'   // 标记为已访问

		for _, dir := range directions {
			if dfs(x+dir.dx, y+dir.dy, index+1) {
				return true // 如果找到单词，返回 true
			}
		}

		board[x][y] = temp // 恢复原字符
		return false       // 未找到单词
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true // 如果从任意位置找到单词，返回 true
			}
		}
	}

	return false // 未找到单词
}

// 46 全排列 https://leetcode.cn/problems/permutations/description/
// permute 返回给定整数数组的所有全排列
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
func permute(nums []int) [][]int {
	result := [][]int{}             // 存放所有排列的结果
	path := []int{}                 // 当前排列路径
	used := make([]bool, len(nums)) // 用于标记每个数字是否被使用

	var backtrack func()

	backtrack = func() {
		// 当路径长度等于 nums 的长度时，添加到结果中
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...)) // 复制当前路径
			return
		}

		// 遍历所有数字
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue // 如果数字已被使用，跳过
			}
			// 选择当前数字
			path = append(path, nums[i])
			used[i] = true // 标记为已使用

			// 递归调用
			backtrack()

			// 撤销选择
			path = path[:len(path)-1]
			used[i] = false // 标记为未使用
		}
	}

	backtrack() // 从空路径开始生成排列
	return result
}

// 47 全排列ii https://leetcode.cn/problems/permutations-ii/description/
// permuteUnique 返回给定整数数组的所有不同全排列
// 输入：nums = [1,1,2]
// 输出：[[1,1,2],[1,2,1],[2,1,1]]
func permuteUnique(nums []int) [][]int {
	result := [][]int{}             // 存放所有全排列的结果
	path := []int{}                 // 当前排列路径
	used := make([]bool, len(nums)) // 用于标记每个数字是否被使用

	sort.Ints(nums) // 排序数组，以便处理重复元素

	var backtrack func()

	backtrack = func() {
		// 当路径长度等于 nums 的长度时，添加到结果中
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...)) // 复制当前路径
			return
		}

		// 遍历所有数字
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue // 如果数字已被使用，跳过
			}
			// 处理重复元素，确保同一树层只选择一次相同元素
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue // 跳过重复元素
			}
			// 选择当前数字
			path = append(path, nums[i])
			used[i] = true // 标记为已使用

			// 递归调用
			backtrack()

			// 撤销选择
			path = path[:len(path)-1]
			used[i] = false // 标记为未使用
		}
	}

	backtrack() // 从空路径开始生成排列
	return result
}

// 131 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/description/
func partition(s string) [][]string {
	result := [][]string{} // 存放所有分割结果
	path := []string{}     // 当前分割路径

	var backtrack func(start int)

	backtrack = func(start int) {
		// 如果起始位置到字符串末尾，说明找到一个分割
		if start == len(s) {
			result = append(result, append([]string{}, path...)) // 复制当前路径
			return
		}

		// 遍历可能的分割位置
		for end := start + 1; end <= len(s); end++ {
			substr := s[start:end]    // 当前子串
			if isPalindrome(substr) { // 检查是否为回文串
				path = append(path, substr) // 选择当前子串
				backtrack(end)              // 递归调用，继续分割后续部分
				path = path[:len(path)-1]   // 撤销选择
			}
		}
	}

	backtrack(0) // 从起始位置开始分割
	return result
}

// 辅助函数，检查字符串是否为回文串
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false // 如果不相等，则不是回文串
		}
		left++
		right--
	}
	return true // 如果所有字符都相等，则是回文串
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	fmt.Println(subsets([]int{1, 2, 3}))
	permute := permute([]int{1, 2, 3})
	fmt.Println(permute)
	partition := partition("aab")
	fmt.Println(partition)
}
