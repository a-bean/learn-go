package main

import "fmt"

// 26. https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
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

// 283. https://leetcode.cn/problems/move-zeroes/description/
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

// 88. https://leetcode.cn/problems/merge-sorted-array/description/
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

func main() {
	fmt.Print(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	moveZeroes([]int{1, 1, 1, 0, 2, 2, 3})
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
