package main

import (
	"fmt"
	"math"
)

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

// 35 搜索插入位置 https://leetcode.cn/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left // 如果没有找到，返回插入位置
}

// 33 搜索旋转排序数组 https://leetcode.cn/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // 左半部分有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // 在左半部分继续查找
			} else {
				left = mid + 1 // 在右半部分继续查找
			}
		} else { // 右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // 在右半部分继续查找
			} else {
				right = mid - 1 // 在左半部分继续查找
			}
		}
	}
	return -1 // 如果没有找到，返回 -1
}

// 74. 搜索二维矩阵 https://leetcode.cn/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)>>1
		midValue := matrix[mid/cols][mid%cols]

		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 4. 寻找两个正序数组的中位数 https://leetcode.cn/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	}
	midIndex1, midIndex2 := totalLength/2-1, totalLength/2
	return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMedianSortedArrays1(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a // 保证下面的 i 可以从 0 开始枚举， 确保 a 是较短的数组，以减少二分查找的复杂度
	}

	m, n := len(a), len(b)
	a = append([]int{math.MinInt}, append(a, math.MaxInt)...)
	b = append([]int{math.MinInt}, append(b, math.MaxInt)...)

	// 枚举 nums1 有 i 个数在第一组
	// 那么 nums2 有 j = (m+n+1)/2 - i 个数在第一组
	i, j := 0, (m+n+1)/2
	for {
		if a[i] <= b[j+1] && a[i+1] > b[j] { // 写 >= 也可以
			max1 := max(a[i], b[j])     // 第一组的最大值
			min2 := min(a[i+1], b[j+1]) // 第二组的最小值
			if (m+n)%2 > 0 {
				return float64(max1)
			}
			return float64(max1+min2) / 2
		}
		i++ // 继续枚举
		j--
	}
}

func main() {
	findMin([]int{4, 5, 6, 1, 2, 3})
	findMin2([]int{4, 5, 6, 1, 2, 3})
	searchRange([]int{5, 7, 7, 8, 8, 10}, 8)

	searchLastEqualElement([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(mySqrt(9), mySqrt(8))
	findPeakElement([]int{1, 2, 3, 1})
	findPeakElement2([]int{1, 2, 3, 1})

	splitArray([]int{7, 2, 5, 10, 8}, 2)
	findMedianSortedArrays1([]int{1, 2}, []int{3, 4})

}
