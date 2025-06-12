package main

import (
	"fmt"
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

// 152 乘积最大数组 https://leetcode-cn.com/problems/maximum-product-subarray/
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
// 01背包问题变种
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

// 518 零钱兑换 II https://leetcode-cn.com/problems/coin-change-2/
// 组合数问题，计算不同的硬币组合数
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // 组合数初始化为1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}

	return dp[amount]
}

// 32 最长有效括号 https://leetcode-cn.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 0
	stack := []int{-1} // 初始化栈，-1 用于处理边界情况

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i) // 遇到左括号，入栈
		} else {
			stack = stack[:len(stack)-1] // 遇到右括号，出栈
			if len(stack) == 0 {
				stack = append(stack, i) // 如果栈空了，入栈当前索引
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1]) // 更新最大长度
			}
		}
	}

	return maxLen
}

// dp解法
func longestValidParenthesesDP(s string) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	dp := make([]int, n)
	maxLen := 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {

			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2 // 匹配到一对括号
				} else {
					dp[i] = 2 // 匹配到一对括号
				}

			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {

				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2] // 匹配到嵌套括号
				} else {
					dp[i] = dp[i-1] + 2 // 匹配到嵌套括号
				}

			}
			maxLen = max(maxLen, dp[i])
		}
	}

	return maxLen
}

// 62 不同路径 https://leetcode-cn.com/problems/unique-paths/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化第一行和第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 填充 dp 数组
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 64 最小路径和 https://leetcode-cn.com/problems/minimum-path-sum/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0] // 起点

	// 初始化第一行和第一列
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 填充 dp 数组
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

// 5 最长回文子串 https://leetcode-cn.com/problems/longest-palindromic-substring/
// longestPalindrome 查找字符串中最长的回文子串
// 使用中心扩展法：对每个可能的中心点，向两边扩展检查回文
// s: 输入字符串
// 返回: 最长回文子串
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0 // 记录最长回文子串的起始和结束位置

	// 遍历每个可能的中心点
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   // 以单个字符为中心（奇数长度回文）
		len2 := expandAroundCenter(s, i, i+1) // 以两个字符之间为中心（偶数长度回文）
		maxLen := max(len1, len2)             // 取两种情况的较大值

		// 如果找到更长的回文子串，更新起始和结束位置
		if maxLen > end-start {
			start = i - (maxLen-1)/2 // 计算回文串的起始位置
			end = i + maxLen/2       // 计算回文串的结束位置
		}
	}

	return s[start : end+1] // 返回最长回文子串
}

// expandAroundCenter 从中心向两边扩展检查回文
// s: 原字符串
// left: 左边界起始位置
// right: 右边界起始位置
// 返回: 以该中心点扩展得到的回文串的长度
func expandAroundCenter(s string, left int, right int) int {
	// 当左右指针都在有效范围内且对应字符相等时，继续扩展
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--  // 向左扩展
		right++ // 向右扩展
	}
	// 返回回文长度：right-left-1
	// 因为最后一次循环会多执行一次left--和right++
	return right - left - 1
}

// dp解法
// longestPalindromeDP 使用动态规划方法查找最长回文子串
// s: 输入字符串
// 返回: 最长回文子串
func longestPalindromeDP(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	// dp[i][j] 表示从索引i到j的子串是否为回文
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1 // 记录最长回文子串的起始位置和长度

	// 初始化：所有长度为1的子串都是回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 检查长度为2的子串
	// 如果相邻字符相同，则形成回文
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// 检查长度大于2的子串
	// 状态转移方程：dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1 // 子串的结束位置
			// 当前子串是回文的条件：
			// 1. 首尾字符相同
			// 2. 去掉首尾后的子串也是回文
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > maxLength {
					start = i
					maxLength = length
				}
			}
		}
	}

	return s[start : start+maxLength]
}

// 121 	买卖股票的最佳时机 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	minPrice := prices[0] // 初始化最小价格为第一个价格
	for i := 1; i < len(prices); i++ {
		// 计算当前价格与之前的最小价格的差值
		dp[i] = max(dp[i-1], prices[i]-minPrice)
		// 更新最小价格
		minPrice = min(minPrice, prices[i])
	}
	return dp[len(prices)-1] // 返回最大利润
}

// 122 买卖股票的最佳时机 II https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfitII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			// 如果当前价格大于前一个价格，则可以卖出
			dp[i] = dp[i-1] + (prices[i] - prices[i-1])
		} else {
			// 否则保持之前的利润
			dp[i] = dp[i-1]
		}
		maxProfit = max(maxProfit, dp[i])
	}
	return maxProfit
}

// 123 买卖股票的最佳时机 III https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
func maxProfitIII(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	if n < 2 {
		return 0
	}

	// dp[i][j] 表示在第 i 天，最多进行 j 次交易时的最大利润
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 3) // 最多两次交易
	}

	for j := 1; j <= 2; j++ {
		maxDiff := -prices[0] // 初始化最大差值
		for i := 1; i < n; i++ {
			dp[i][j] = max(dp[i-1][j], prices[i]+maxDiff) // 当前利润与之前的利润比较
			maxDiff = max(maxDiff, dp[i][j-1]-prices[i])  // 更新最大差值
		}
	}

	return dp[n-1][2] // 返回最多两次交易的最大利润
}

// 188 买卖股票的最佳时机 IV https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
func maxProfitIV(k int, prices []int) int {
	if len(prices) == 0 || k <= 0 {
		return 0
	}

	n := len(prices)
	if k >= n/2 { // 如果交易次数大于等于天数的一半，直接使用贪心算法
		maxProfit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxProfit += prices[i] - prices[i-1]
			}
		}
		return maxProfit
	}

	// dp[i][j] 表示在第 i 天，最多进行 j 次交易时的最大利润
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	for j := 1; j <= k; j++ {
		maxDiff := -prices[0] // 初始化最大差值
		for i := 1; i < n; i++ {
			// dp[i-1][j]: 第 i 天不进行交易，沿用之前的利润
			// prices[i]+maxDiff: 第 i 天进行卖出操作
			// 										maxDiff 代表前面某一天买入的最大收益可能性
			// 										prices[i] 是当天的卖出价格
			dp[i][j] = max(dp[i-1][j], prices[i]+maxDiff)

			//dp[i][j-1]: 进行了 j-1 次交易到第 i 天的最大利润
			// -prices[i]: 在第 i 天买入的成本
			maxDiff = max(maxDiff, dp[i][j-1]-prices[i]) // 更新最大差值
		}
	}

	return dp[n-1][k] // 返回最多 k 次交易的最大利润
}

// 714 买卖股票的最佳时机含手续费 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
func maxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2) // dp[i][0] 表示第 i 天不持有股票的最大利润，dp[i][1] 表示第 i 天持有股票的最大利润
	}

	dp[0][0] = 0                // 第一天不持有股票，利润为0
	dp[0][1] = -prices[0] - fee // 第一天持有股票，利润为负的买入价格加手续费

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])     // 今天不持有股票，可以是昨天就不持有，或者今天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee) // 今天持有股票，可以是昨天就持有，或者今天买入
	}

	return dp[n-1][0] // 返回最后一天不持有股票的最大利润
}

// 309 买卖股票的最佳时机含冷冻期 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfitWithCooldown(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2) // dp[i][0] 表示第 i 天不持有股票的最大利润，dp[i][1] 表示第 i 天持有股票的最大利润
	}

	dp[0][0] = 0          // 第一天不持有股票，利润为0
	dp[0][1] = -prices[0] // 第一天持有股票，利润为负的买入价格

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) // 今天不持有股票，可以是昨天就不持有，或者今天卖出
		if i > 1 {
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) // 今天持有股票，可以是昨天就持有，或者今天买入（考虑冷冻期）
		} else {
			dp[i][1] = max(dp[i-1][1], -prices[i]) // 第一天或第二天没有冷冻期限制
		}
	}

	return dp[n-1][0] // 返回最后一天不持有股票的最大利润
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

// 213 打家劫舍 II https://leetcode-cn.com/problems/house-robber-ii/
func robII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// 分两种情况：不偷第一个房子或不偷最后一个房子
	return max(rob(nums[:len(nums)-1]), rob(nums[1:]))
}

// 72 编辑距离 https://leetcode-cn.com/problems/edit-distance/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n // 如果第一个字符串为空，返回第二个字符串的长度
	}
	if n == 0 {
		return m // 如果第二个字符串为空，返回第一个字符串的长度
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化第一行和第一列
	for i := 0; i <= m; i++ {
		dp[i][0] = i // 删除操作
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // 插入操作
	}

	// 填充 dp 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] // 字符相同，不需要操作
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1)) // 删除、插入、替换操作
			}
		}
	}

	return dp[m][n]
}

// 01背包问题 https://leetcode-cn.com/problems/partition-equal-subset-sum/

func main() {
	// 示例用法
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	coinChange([]int{5, 2, 1}, amount)
	fmt.Println("最少硬币数:", result) // 输出: 最少硬币数: 3

	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	result1 := uniquePathsWithObstacles(obstacleGrid)
	fmt.Println("不同路径数:", result1) // 输出: 不同路径数: 2

	longestCommonSubsequence(`abcde`, `ace`)
	lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	length, sequence := lengthOfLISWithSequence([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Println("最长递增子序列长度:", length) // 输出: 最长递增子序列长度: 4
	fmt.Println("最长递增子序列:", sequence) // 输出: 最长递增子序列: [2 3 7 101]

	wordBreak(`catsandog`, []string{"cats", "dog", "sand", "and", "cat"})
	maxProduct([]int{2, 3, -2, 4})
	canPartition([]int{1, 5, 11, 5})
	canPartition([]int{1, 2, 3, 5})
	fmt.Print("最长有效括号长度: ")
	fmt.Println(longestValidParentheses(`(())(())`)) // 输出: 4
	minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	})

	fmt.Println("最长回文子串:", longestPalindrome(`babad`))  // 输出: bab 或者 aba
	fmt.Println("最长回文子串:", longestPalindromeDP(`cbbd`)) // 输出: bb

	// minDistance(`intention`, `execution`)
	minDistance(`horse`, `ros`)

	numSquares(13)

	a := "catsandog"
	fmt.Println(a[3:6])
}
