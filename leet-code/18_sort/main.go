package main

import (
	"sort"
)

// 1122 数组的相对排序 https://leetcode.cn/problems/relative-sort-array/
// 解法一 桶排序，时间复杂度 O(n^2)
func relativeSortArray(A, B []int) []int {
	// 按桶映射
	count := [1001]int{}
	for _, a := range A {
		count[a]++
	}

	res := make([]int, 0, len(A))
	// 按B顺序输出
	for _, b := range B {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}
	// 按顺序输出B没有的元素
	for i := 0; i < 1001; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}
	return res
}

// 解法二 模拟，时间复杂度 O(n^2)
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	leftover, m, res := []int{}, make(map[int]int), []int{}
	for _, v := range arr1 {
		m[v]++
	}

	for _, s := range arr2 {
		count := m[s]
		for i := 0; i < count; i++ {
			res = append(res, s)
		}
		m[s] = 0
	}

	for v, count := range m {
		for i := 0; i < count; i++ {
			leftover = append(leftover, v)
		}
	}

	sort.Ints(leftover)
	res = append(res, leftover...)
	return res
}

// 56 合并区间 https://leetcode.cn/problems/merge-intervals/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// merge 函数合并重叠的区间
func merge(intervals [][]int) [][]int {
	// 如果输入的区间为空，直接返回空切片
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 对区间进行排序，按起始值升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}                   // 初始化结果，添加第一个区间
	result = append(result, intervals[0]) // 添加第一个区间

	// 遍历区间，开始合并
	for i := 1; i < len(intervals); i++ {
		// 获取最后一个合并后的区间
		last := result[len(result)-1]

		// 检查当前区间与最后一个合并后区间是否重叠
		if intervals[i][0] <= last[1] { // 有重叠
			last[1] = max(last[1], intervals[i][1]) // 更新结束值
		} else {
			// 没有重叠，添加当前区间
			result = append(result, intervals[i])
		}
	}

	return result // 返回合并后的区间
}

// 493 翻转对 https://leetcode.cn/problems/reverse-pairs/
// reversePairs 函数计算数组中逆序对的数量
func reversePairs(nums []int) int {
	if len(nums) < 2 { // 如果数组长度小于2，直接返回0
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1) // 调用归并排序并返回逆序对的数量
}

// mergeSort 函数使用归并排序算法计算逆序对的数量
func mergeSort(nums []int, left, right int) int {
	if left >= right { // 如果左索引大于或等于右索引，返回0
		return 0
	}
	mid := left + (right-left)/2                                        // 计算中间索引
	count := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right) // 递归计算左右部分的逆序对数量

	j := mid + 1 // 初始化右侧数组的指针
	// 计算逆序对
	for i := left; i <= mid; i++ {
		// 找到 nums[i] > 2 * nums[j] 的所有 j
		for j <= right && nums[i] > 2*nums[j] {
			j++ // 移动 j 指针
		}
		count += j - (mid + 1) // 统计逆序对数量
	}

	merge1(nums, left, mid, right) // 合并两个已排序的子数组
	return count                   // 返回逆序对的数量
}

// merge1 函数合并两个已排序的子数组
func merge1(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1) // 创建临时数组用于存储合并结果
	i, j, k := left, mid+1, 0         // 初始化指针

	// 合并两个子数组
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temp[k] = nums[i] // 将较小的元素放入临时数组
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}

	// 处理左侧剩余元素
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}

	// 处理右侧剩余元素
	for j <= right {
		temp[k] = nums[j]
		j++
		k++
	}

	// 将合并后的结果复制回原数组
	for i := 0; i < len(temp); i++ {
		nums[left+i] = temp[i]
	}
}

func main() {
	relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	relativeSortArray1([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	reversePairs([]int{7, 5, 6, 4})

}
