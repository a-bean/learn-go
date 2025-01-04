package main

import (
	"fmt"
	"math"
	"sort"
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

// 15. 三数之和 https://leetcode.cn/problems/3sum/description/
// 解法一: 最优解，双指针 + 排序
func threeSum1(nums []int) [][]int {
	// 特判
	res, length, L, R, sum := make([][]int, 0), len(nums), 0, 0, 0
	if length < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)

	for i := 0; i < length; i++ {
		// 如果遍历的起始元素大于0，就直接退出
		// 原因，此时数组为有序的数组，最小的数都大于0了，三数之和肯定大于0
		if nums[i] > 0 {
			break
		}
		// 去重，当起始的值等于前一个元素，那么得到的结果将会和前一次相同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L, R = i+1, length-1
		for L < R {
			sum = nums[L] + nums[i] + nums[R]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum > 0 {
				R--
			} else if sum < 0 {
				L++
			}
		}
	}
	return res
}

// 解法二
func threeSum2(nums []int) [][]int {
	var res [][]int
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		// 三个0的情况
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// 16. 最接近的三数之和 https://leetcode.cn/problems/3sum-closest/description/
func threeSumClosest(nums []int, target int) int {
	length, closestSum, minDiff := len(nums), 0, math.MaxInt16
	if length > 2 {
		sort.Ints(nums)
		for i := 0; i < length-2; i++ {
			// 判重
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			L, R := i+1, length-1
			for L < R {
				sum := nums[i] + nums[L] + nums[R]
				if abs(sum-target) < minDiff {
					closestSum, minDiff = sum, abs(sum-target)
				}

				if sum == target {
					return closestSum
				} else if sum > target {
					R--
				} else {
					L++
				}
			}
		}
	}
	return closestSum
}

// 18. 四数之和 https://leetcode.cn/problems/4sum/description/
func fourSum(nums []int, target int) [][]int {
	var quadruplets [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 第i个与第i-1个相等 或者 第i个与最大的三个数相加 还小于target,就跳过
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return quadruplets
}

// 解法二 kSum
// 解法一 双指针
func fourSum1(nums []int, target int) [][]int {
	res, cur := make([][]int, 0), make([]int, 0)
	sort.Ints(nums)
	kSum(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}

func kSum(nums []int, left, right int, target int, k int, cur []int, res *[][]int) {
	if right-left+1 < k || k < 2 || target < nums[left]*k || target > nums[right]*k {
		return
	}
	if k == 2 {
		// 2 sum
		twoSum1(nums, left, right, target, cur, res)
	} else {
		for i := left; i < len(nums); i++ {
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]int, len(cur))
				copy(next, cur)
				next = append(next, nums[i])
				kSum(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}

}

func twoSum1(nums []int, left, right int, target int, cur []int, res *[][]int) {
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			cur = append(cur, nums[left], nums[right])
			temp := make([]int, len(cur))
			copy(temp, cur)
			*res = append(*res, temp)
			// reset cur to previous state
			cur = cur[:len(cur)-2]
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}

// 75. 排序 https://leetcode.cn/problems/sort-colors/description/
func sortColors(nums []int) {
	zero, one := 0, 0
	for index, value := range nums {
		nums[index] = 2
		if value <= 1 {
			nums[one] = 1
			one++
		}
		if value == 0 {
			nums[zero] = 0
			zero++
		}
	}
}

func main() {
	fmt.Print(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	moveZeroes([]int{1, 1, 1, 0, 2, 2, 3})
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	fmt.Print(twoSum([]int{1, 2, 3, 4}, 5))
	threeSum1([]int{-1, 0, 1, 2, -1, -4})
	threeSum2([]int{-1, 0, 1, 2, -1, -4})
	fmt.Print(threeSumClosest([]int{1, 1, 1, 2, 3}, 3))
	fmt.Print(fourSum([]int{1, 1, 1, 2, 3}, 3))
	fmt.Print(fourSum1([]int{1, 1, 1, 2, 3}, 3))
	sortColors([]int{1, 1, 1, 0, 0, 0, 2})
}
