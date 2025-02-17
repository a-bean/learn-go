package main

import (
	"sort"
)

// 860 柠檬水找零 https://leetcode.cn/problems/lemonade-change/description/
// lemonadeChange 函数判断是否能够找零给顾客
func lemonadeChange(bills []int) bool {
	// 创建一个地图来记录当前拥有的硬币数量
	coins := map[int]int{5: 0, 10: 0, 20: 0}

	// exchange 函数用于找零
	var exchange = func(amount int) bool {
		// 遍历硬币面额，从大到小(贪心所在,优先使用大面额的硬币)
		for _, coin := range []int{20, 10, 5} {
			// 当还有硬币且找零金额大于等于当前硬币面额时
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
func lemonadeChange1(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
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
// 决策范围扩展

// 45 跳跃游戏 https://leetcode.cn/problems/jump-game/
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

// 1665 完成所有任务的最少初始能量 https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/
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

func main() {
	lemonadeChange([]int{5, 5, 5, 10, 20})
	lemonadeChange1([]int{5, 5, 5, 10, 20})

	canJump([]int{2, 3, 1, 1, 4})
	canJump([]int{3, 2, 1, 0, 4})

	minimumEffort([][]int{{1, 2}, {2, 4}, {4, 8}})
	minimumEffort([][]int{{1, 3}, {5, 4}, {4, 6}})

	findContentChildren([]int{1, 2, 3}, []int{1, 1})
	findContentChildren([]int{1, 2}, []int{1, 2, 3})

}
