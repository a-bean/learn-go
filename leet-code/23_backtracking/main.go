package main

// 46 全排列 https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	var res [][]int

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums)-1 {
			// 复制当前排列到结果中
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start] // 交换
			backtrack(start + 1)                        // 递归调用
			nums[start], nums[i] = nums[i], nums[start] // 回溯
		}

	}

	backtrack(0)
	return res
}

func main() {
	permute([]int{1, 2, 3})
}
