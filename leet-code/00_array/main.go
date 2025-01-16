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

func main() {
	fmt.Print(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	moveZeroes([]int{1, 1, 1, 0, 2, 2, 3})
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	fmt.Print(twoSum([]int{1, 2, 3, 4}, 5))

}
