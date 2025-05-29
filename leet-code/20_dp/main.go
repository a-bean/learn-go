package main

import (
	"fmt"
	"sort"
)

// 322 零钱兑换问题 https://leetcode-cn.com/problems/coin-change/
// 动态规划解法
// 1. 定义状态：dp[i] 表示金额 i 的最小硬币数
// 2. 状态转移方程：dp[i] = min(dp[i - coin] + 1) for each coin in coins
// 3. 初始化状态：dp[0] = 0, dp[i] = inf for i > 0
// 4. 计算顺序：从小到大计算 dp 数组
// 5. 返回结果：dp[amount] 如果 dp[amount] == inf 则返回 -1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // 初始化为一个大数
	}
	dp[0] = 0 // 零元需要零个硬币

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1 // 无法凑成该金额
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 63 路径计数II https://leetcode-cn.com/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 1 {
		return 0 // 起点被障碍物阻挡
	}
	dp[0][0] = 1 // 起点

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0 // 遇到障碍物
			} else {
				if i > 0 {
					dp[i][j] += dp[i-1][j] // 从上方来
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1] // 从左方来
				}
			}
		}
	}

	return dp[m-1][n-1]
}

// 1143 最长公共子序列 https://leetcode-cn.com/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 300 最长递增子序列 https://leetcode-cn.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = 1 // 每个元素至少可以形成一个长度为1的递增子序列
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxLength := 0
	for _, length := range dp {
		maxLength = max(maxLength, length)
	}

	return maxLength
}

// 输入最长递增子序列结果
func lengthOfLISWithSequence(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nil
	}

	n := len(nums)
	dp := make([]int, n)   // dp[i] 表示以 nums[i] 结尾的最长递增子序列长度
	prev := make([]int, n) // prev[i] 记录前驱节点的索引

	// 初始化
	for i := range dp {
		dp[i] = 1
		prev[i] = -1 // -1 表示没有前驱节点
	}

	// 记录最大长度及其结束位置
	maxLen, maxIndex := 1, 0

	// 动态规划过程
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j // 记录前驱节点
			}
		}
		// 更新最大长度和对应的结束位置
		if dp[i] > maxLen {
			maxLen = dp[i]
			maxIndex = i
		}
	}

	// 构建最长递增子序列
	sequence := make([]int, maxLen)
	for i := maxLen - 1; i >= 0; i-- {
		sequence[i] = nums[maxIndex]
		maxIndex = prev[maxIndex]
	}

	return maxLen, sequence
}

// 53 最大子序和 https://leetcode-cn.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0] // 初始化第一个元素

	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i]) // 如果当前元素大于当前和，则重新开始
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

// 153 乘积最大数组 https://leetcode-cn.com/problems/maximum-product-subarray/
// maxProduct 计算数组中连续子数组的最大乘积
// 由于负数的存在，需要同时维护最大值和最小值
// nums: 输入的整数数组
// 返回: 连续子数组的最大乘积
func maxProduct(nums []int) int {
	// 初始化结果、最大值和最小值都为第一个元素
	res, maxF, minF := nums[0], nums[0], nums[0]

	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 保存当前的最大值和最小值，因为计算新的最大最小值时会用到
		mx, mn := maxF, minF

		// 计算新的最大值：
		// 1. 当前数和之前最大值的乘积
		// 2. 当前数本身
		// 3. 当前数和之前最小值的乘积（处理负数情况）
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))

		// 计算新的最小值：同样考虑三种情况
		minF = min(mx*nums[i], min(nums[i], mn*nums[i]))

		// 处理整数溢出的情况
		if minF < (-1 << 31) {
			minF = nums[i]
		}

		// 更新全局最大乘积
		res = max(res, maxF)
	}
	return res
}

func maxProduct1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := nums[0]
	minProd, curProd := 1, 1

	for _, num := range nums {
		if num < 0 {
			minProd, curProd = curProd, minProd // 交换最大和最小乘积
		}
		curProd = max(num, curProd*num) // 当前最大乘积
		minProd = min(num, minProd*num) // 当前最小乘积

		maxProd = max(maxProd, curProd) // 更新全局最大乘积
	}

	return maxProd
}

// 70 爬楼梯问题 https://leetcode-cn.com/problems/climbing-stairs/
func climbingStairs(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 118 杨辉三角 https://leetcode-cn.com/problems/pascals-triangle/
func generatePascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1 // 每行的首尾元素为1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j] // 当前元素等于上一行的两个元素之和
		}
	}
	return triangle
}

// 198 打家劫舍 https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

// 279 完全平方数 https://leetcode-cn.com/problems/perfect-squares/
func numSquares(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // 最坏情况，每个数都可以表示为1的平方和
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

// 139 单词拆分 https://leetcode-cn.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true // 空字符串可以被拆分

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// 三角形的最小路径和 https://leetcode-cn.com/problems/triangle/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	n := len(triangle)
	dp := make([]int, n)
	copy(dp, triangle[n-1]) // 初始化为最后一行

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}

// 416 分割等和子集 https://leetcode-cn.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false // 如果总和是奇数，无法分割成两部分
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true // 零总和可以通过不选任何元素实现

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[target]
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	if nums[0] > 0 {
		return [][]int{}
	}

	left, right := 1, len(nums)-1
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				res = append(res, []int{i, left, right})
			}
		}

	}
	return res
}

func main() {
	// 示例用法
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	fmt.Println("最少硬币数:", result) // 输出: 最少硬币数: 3

	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	result1 := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println("不同路径数:", result1) // 输出: 不同路径数: 2

	longestCommonSubsequence(`abcde`, `ace`)
	lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})
	length, sequence := lengthOfLISWithSequence([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Println("最长递增子序列长度:", length) // 输出: 最长递增子序列长度: 4
	fmt.Println("最长递增子序列:", sequence) // 输出: 最长递增子序列: [2 3 7 101]

	wordBreak(`catsandog`, []string{"cats", "dog", "sand", "and", "cat"})
	maxProduct([]int{2, 3, -2, 4})
	canPartition([]int{1, 5, 11, 5})
	canPartition([]int{1, 2, 3, 5})
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
