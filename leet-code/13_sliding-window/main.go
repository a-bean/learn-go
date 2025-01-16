package main

import (
	"container/heap"
	"fmt"
)

// 567 字符串的排列  https://leetcode.cn/problems/permutation-in-string/description/
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}

	var freq [26]rune
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}

	left, right, count := 0, 0, len(s1)

	for right < len(s2) {

		if freq[s2[right]-'a'] >= 1 {
			count--
		}

		freq[s2[right]-'a']--
		right++

		if count == 0 {
			return true
		}

		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}

	}
	return false
}

// 763 划分字母区间  https://leetcode.cn/problems/partition-labels/description/
// 输入：s = "ababcbaca defegdehijhklij"
// 输出:[9,7,8]
func partitionLabels(S string) []int {
	var lastIndexOf [26]int // 用于记录每个字母最后出现的索引

	// 遍历字符串 S，记录每个字母最后出现的位置
	for i, v := range S {
		lastIndexOf[v-'a'] = i // 将字母 v 的最后索引存储在 lastIndexOf 数组中
	}

	var arr []int // 用于存储每个区间的长度
	// start 表示当前区间的起始位置，end 表示当前区间的结束位置
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = lastIndexOf[S[start]-'a'] // 设置当前区间的结束位置为当前字母的最后出现位置

		// 扩展当前区间的结束位置
		for i := start; i < end; i++ {
			// 如果当前字母的最后出现位置在 end 之后，更新 end
			if end < lastIndexOf[S[i]-'a'] {
				end = lastIndexOf[S[i]-'a']
			}
		}
		// 将当前区间的长度添加到结果数组中
		arr = append(arr, end-start+1)
	}
	return arr // 返回所有区间的长度
}

// TODO: 还没研究
// 480 滑动窗口中位数  https://leetcode.cn/problems/sliding-window-median/description/

// 定义最大堆和最小堆
type MaxHeap []int
type MinHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// findMedian 在最大堆和最小堆中找到中位数
func findMedian(maxHeap *MaxHeap, minHeap *MinHeap) float64 {
	if maxHeap.Len() > minHeap.Len() {
		return float64((*maxHeap)[0])
	}
	return float64((*maxHeap)[0]+(*minHeap)[0]) / 2.0
}

// medianSlidingWindow 计算滑动窗口的中位数
func medianSlidingWindow(nums []int, k int) []float64 {
	if k == 0 {
		return []float64{}
	}

	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	result := []float64{}

	for i := 0; i < len(nums); i++ {
		// 添加新元素
		if maxHeap.Len() == 0 || nums[i] <= (*maxHeap)[0] {
			heap.Push(maxHeap, nums[i])
		} else {
			heap.Push(minHeap, nums[i])
		}

		// 平衡两个堆
		if maxHeap.Len() > minHeap.Len()+1 {
			heap.Push(minHeap, heap.Pop(maxHeap))
		} else if minHeap.Len() > maxHeap.Len() {
			heap.Push(maxHeap, heap.Pop(minHeap))
		}

		// 当窗口大小达到 k 时，计算中位数
		if i >= k-1 {
			result = append(result, findMedian(maxHeap, minHeap))

			// 移除滑动窗口的元素
			toRemove := nums[i-k+1]
			if toRemove <= (*maxHeap)[0] {
				// 从最大堆中移除
				for j := 0; j < maxHeap.Len(); j++ {
					if (*maxHeap)[j] == toRemove {
						(*maxHeap)[j] = (*maxHeap)[maxHeap.Len()-1]
						*maxHeap = (*maxHeap)[:maxHeap.Len()-1]
						heap.Init(maxHeap)
						break
					}
				}
			} else {
				// 从最小堆中移除
				for j := 0; j < minHeap.Len(); j++ {
					if (*minHeap)[j] == toRemove {
						(*minHeap)[j] = (*minHeap)[minHeap.Len()-1]
						*minHeap = (*minHeap)[:minHeap.Len()-1]
						heap.Init(minHeap)
						break
					}
				}
			}

			// 重新平衡堆
			if maxHeap.Len() > minHeap.Len()+1 {
				heap.Push(minHeap, heap.Pop(maxHeap))
			} else if minHeap.Len() > maxHeap.Len() {
				heap.Push(maxHeap, heap.Pop(minHeap))
			}
		}
	}

	return result
}

// 978 最长公共子序列 https://leetcode.cn/problems/longest-turbulent-subarray/description/

func maxTurbulenceSize(arr []int) int {
	var maxLength int
	if len(arr) == 2 && arr[0] != arr[1] {
		maxLength = 2
	} else {
		maxLength = 1
	}
	left := 0
	for right := 2; right < len(arr); right++ {
		if arr[right] == arr[right-1] {
			left = right
		} else if (arr[right]-arr[right-1])^(arr[right-1]-arr[right-2]) >= 0 {
			left = right - 1
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	checkInclusion("ab", "eidbaooo")
	partitionLabels("ababcbacadefegdehijhklij")

	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := medianSlidingWindow(nums, k)
	fmt.Println(result) // 输出滑动窗口的中位数

	maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9})
}
