package main

import "fmt"

// 153 寻找旋转排序数组中的最小值 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {

		if nums[low] < nums[high] {
			return nums[low]
		}

		mid := low + (high-low)>>1

		if nums[mid] >= nums[low] {
			low = mid + 1
		} else {
			high = mid
		}

	}
	return nums[low]
}

// 154 寻找旋转排序数组中的最小值2 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii
func findMin2(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)>>1

		if nums[mid] > nums[high] {
			low = mid + 1
		} else if nums[mid] < nums[high] {
			high = mid
		} else {
			high--
		}
	}
	return nums[low]
}

// 34 在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array
func searchRange(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}
func searchFirstEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			if (mid == 0) || (nums[mid-1] != target) {
				return mid
			}
			right = mid
		}
	}
	return -1
}

func searchLastEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) {
				return mid
			}
			left = mid
		}
	}
	return -1
}

// 69 x 的平方根 https://leetcode.cn/problems/sqrtx/
func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		mid := left + (right+1-left)>>1
		if mid <= x/mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

// 162 寻找峰值 https://leetcode.cn/problems/find-peak-element/
// 二分查找
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 三分查找
func findPeakElement2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		third := (right - left) / 3
		mid1 := left + third
		mid2 := right - third
		if nums[mid1] < nums[mid2] {
			left = mid1 + 1
		} else {
			right = mid2 - 1
		}
	}
	return left
}

// 410 分割数组的最大值 https://leetcode.cn/problems/split-array-largest-sum/
// splitArray 将数组 nums 分割成 m 个子数组，使得这些子数组的最大和最小。
// 使用二分查找来找到最小的最大子数组和。
func splitArray(nums []int, m int) int {
	low, high := 0, 0

	// 初始化 low 和 high
	// low 是数组中的最大值，因为每个子数组的和至少要大于等于这个值
	// high 是数组所有元素的和，因为这是所有元素在一个子数组中的情况
	for _, num := range nums {
		if low < num {
			low = num
		}
		high += num
	}

	// 使用二分查找来找到最小的最大子数组和
	for low < high {
		mid := low + (high-low)>>1
		// 检查是否可以将数组分割成 m 个或更少的子数组，每个子数组的和不超过 mid
		if check(nums, m, mid) {
			high = mid // 如果可以，尝试更小的最大子数组和
		} else {
			low = mid + 1 // 如果不可以，增加最大子数组和
		}
	}
	return low // 返回最小的最大子数组和
}

// check 检查是否可以将数组 nums 分割成 m 个或更少的子数组，每个子数组的和不超过 target
func check(nums []int, m int, target int) bool {
	sum, count := 0, 1 // 初始化当前子数组的和为 0，子数组计数为 1
	for _, num := range nums {
		if sum+num > target {
			// 如果当前子数组的和加上 num 超过了 target，开始一个新的子数组
			sum = num
			count++
			if count > m {
				// 如果子数组的数量超过了 m，返回 false
				return false
			}
		} else {
			// 否则，将 num 加入当前子数组
			sum += num
		}
	}
	return true // 如果子数组的数量不超过 m，返回 true
}

// 287 寻找重复数 https://leetcode.cn/problems/find-the-duplicate-number/

func main() {
	findMin([]int{4, 5, 6, 1, 2, 3})
	findMin2([]int{4, 5, 6, 1, 2, 3})
	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)

	searchLastEqualElement([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(mySqrt(9), mySqrt(8))
	findPeakElement([]int{1, 2, 3, 1})
	findPeakElement2([]int{1, 2, 3, 1})

	splitArray([]int{7, 2, 5, 10, 8}, 2)
}
