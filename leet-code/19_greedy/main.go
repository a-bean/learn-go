package main

import (
	"sort"
)

// 860 柠檬水找零 https://leetcode.cn/problems/lemonade-change/description/
func lemonadeChange(bills []int) bool {
	// 创建一个地图来记录当前拥有的硬币数量
	coins := map[int]int{5: 0, 10: 0, 20: 0}
	var exchange = func(amount int) bool {
		// 遍历硬币面额，成倍数关系，从大到小(贪心所在,优先使用大面额的硬币) 决策包容。
		for _, coin := range []int{20, 10, 5} {
			for coins[coin] > 0 && amount >= coin {
				amount -= coin // 减去硬币面额
				coins[coin]--  // 硬币数量减一
			}
		}
		return amount == 0 // 如果找零金额为0，返回 true
	}

	// 遍历每一笔交易
	for _, bill := range bills {
		coins[bill]++ // 增加当前收到的钞票数量
		// 尝试找零，如果找零失败，返回 false
		if !exchange(bill - 5) {
			return false
		}
	}
	return true // 如果所有交易都能找零，返回 true
}

// 455 分发饼干 https://leetcode.cn/problems/assign-cookies/
// 决策包容性原则：优先满足胃口
func findContentChildren(g []int, s []int) int {
	sort.Ints(g) // 对孩子的胃口进行排序
	sort.Ints(s) // 对饼干的大小进行排序

	childIndex := 0  // 孩子的索引
	cookieIndex := 0 // 饼干的索引

	// 遍历孩子和饼干
	for childIndex < len(g) && cookieIndex < len(s) {
		if s[cookieIndex] >= g[childIndex] { // 如果当前饼干可以满足当前孩子
			childIndex++ // 满足一个孩子，移动到下一个孩子
		}
		cookieIndex++ // 移动到下一个饼干
	}

	return childIndex // 返回满足的孩子数量
}

// 122 买卖股票的最佳时机 II https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit122(prices []int) int {
	profit := 0 // 初始化利润为0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] { // 如果当前价格大于前一个价格
			profit += prices[i] - prices[i-1] // 累加利润
		}
	}
	return profit // 返回总利润
}

// 121 买卖股票的最佳时机 https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit121(prices []int) int {
	max := 0
	min := 10000
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price-min > max {
			max = price - min
		}
	}

	return max

}

// 55 跳跃游戏 https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	maxReach := 0 // 记录当前能达到的最大索引
	for i := 0; i < len(nums); i++ {
		if i > maxReach { // 如果当前索引超出最大可达范围，返回 false
			return false
		}
		// 更新最大可达范围
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
		// 如果已经可以到达最后一个索引，返回 true
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false // 遍历结束后仍无法到达最后一个索引
}

// 45 跳跃游戏 II https://leetcode-cn.com/problems/jump-game-ii/
func jump(nums []int) int {
	end, maxPos, steps := 0, 0, 0 // 初始化终点、最大位置和步数
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i]) // 更新最大位置 maxPos存上一次最大的值，与当前值比较，上一次更大则选择上一次
		if i == end {                   // 如果到达了当前的终点
			end = maxPos // 更新终点为最大位置
			steps++      // 步数加一
		}
	}
	return steps // 返回步数
}

// 1665 完成所有任务的最少初始能量 https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/
// TODO 临项交换
func minimumEffort(tasks [][]int) int {
	// 按照任务的最小努力值和最大努力值的差值进行排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0]-tasks[i][1] < tasks[j][0]-tasks[j][1]
	})

	ans := 0 // 初始化答案，表示所需的最小初始能量
	// 从最后一个任务开始向前遍历
	for i := len(tasks) - 1; i >= 0; i-- {
		// 更新答案为当前任务的最大努力值和当前答案加上任务的最小努力值中的较大值
		ans = max(tasks[i][1], ans+tasks[i][0])
	}
	return ans // 返回计算得到的最小初始能量
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumEffort1(tasks [][]int) int {
	// 按照任务的最小努力值和最大努力值的差值进行排序
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][0]-tasks[i][1] < tasks[j][0]-tasks[j][1]
	})
	totalEnergy := 0     // 所需的总能量
	remainingEnergy := 0 // 剩余能量

	for i := 0; i < len(tasks); i++ {
		// 如果剩余能量不足以开始任务
		if remainingEnergy < tasks[i][1] {
			// 需要补充的能量
			additionalEnergy := tasks[i][1] - remainingEnergy
			totalEnergy += additionalEnergy
			remainingEnergy = tasks[i][1]
		}
		// 完成任务后的剩余能量
		remainingEnergy -= tasks[i][0]
	}

	return totalEnergy
}

// 763 划分字母区间 https://leetcode.cn/problems/partition-labels/
func partitionLabels(s string) []int {
	lastIndex := make([]int, 26) // 记录每个字母最后出现的位置
	for i, c := range s {
		lastIndex[c-'a'] = i // 更新字母的最后位置
	}

	start, end := 0, 0 // 初始化起始和结束位置
	result := []int{}  // 存储结果

	for i, c := range s {
		end = max(end, lastIndex[c-'a']) // 更新当前区间的结束位置
		if i == end {                    // 如果当前索引等于结束位置
			result = append(result, end-start+1) // 添加区间长度到结果
			start = i + 1                        // 更新起始位置为下一个字符
		}
	}

	return result // 返回划分的区间长度列表
}

func main() {
	lemonadeChange([]int{5, 5, 5, 10, 20})

	canJump([]int{2, 3, 1, 1, 4})
	canJump([]int{3, 2, 1, 0, 4})

	minimumEffort([][]int{{1, 2}, {2, 4}, {4, 8}})
	minimumEffort([][]int{{1, 3}, {5, 4}, {4, 6}})
	minimumEffort([][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}})

	findContentChildren([]int{1, 2, 3}, []int{1, 1})
	findContentChildren([]int{1, 2}, []int{1, 2, 3})

	jump([]int{2, 3, 1, 1, 4})

}
