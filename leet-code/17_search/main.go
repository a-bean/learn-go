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

func main() {
	findMin([]int{4, 5, 6, 1, 2, 3})
	findMin2([]int{4, 5, 6, 1, 2, 3})
	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)

	searchLastEqualElement([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(mySqrt(9), mySqrt(8))

}
