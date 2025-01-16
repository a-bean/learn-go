package main

import (
	"fmt"
	"math"
	"sort"
)

// 11 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water/description/
// 双指针
func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max

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

// 287 寻找重复数 https://leetcode.cn/problems/find-the-duplicate-number/description/
// 输入：nums = [1,3,4,2,2]
// 输出：2
// TODO: 搞不懂这解法
// 想象成链表  使用快慢指针
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast
}

// 解法二 二分搜索
func findDuplicate1(nums []int) int {
	// 初始化二分查找的低、高边界
	low, high := 0, len(nums)-1

	// 进行二分查找，直到低边界小于高边界
	for low < high {
		// 计算中间值，使用位运算提高效率
		mid := low + (high-low)>>1
		count := 0 // 统计小于等于 mid 的元素个数

		// 遍历整个数组，统计有多少个元素小于等于 mid
		for _, num := range nums {
			if num <= mid {
				count++ // 如果当前元素小于等于 mid，计数加一
			}
		}

		if count > mid { // 取左边
			high = mid // 缩小范围，设置高边界为 mid
		} else { // 取右边
			low = mid + 1 // 否则，设置低边界为 mid + 1
		}
	}

	// 返回找到的重复数字，low 和 high 会重合在重复数字的位置
	return low
}

func main() {
	threeSum1([]int{-1, 0, 1, 2, -1, -4})
	threeSum2([]int{-1, 0, 1, 2, -1, -4})
	fmt.Print(threeSumClosest([]int{1, 1, 1, 2, 3}, 3))
	fmt.Print(fourSum([]int{1, 1, 1, 2, 3}, 3))
	fmt.Print(fourSum1([]int{1, 1, 1, 2, 3}, 3))
	sortColors([]int{1, 1, 1, 0, 0, 0, 2})

	nums := []int{1, 2, 4, 3, 3} // 示例数组
	duplicate := findDuplicate(nums)
	duplicate1 := findDuplicate1(nums)
	fmt.Println("找到的重复数字是:", duplicate)
	fmt.Println("找到的重复数字是:", duplicate1)
}
