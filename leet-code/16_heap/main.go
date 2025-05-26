package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 23 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/description/

// 定义最小堆
type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// 初始化最小堆
	h := &MinHeap{}
	heap.Init(h)

	// 将所有链表的头节点加入堆
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	// 创建一个虚拟头节点
	dummy := &ListNode{}
	current := dummy

	// 从堆中取出最小节点，并将其下一个节点加入堆
	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		current.Next = minNode
		current = current.Next

		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
	}

	return dummy.Next
}

// 239 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/description/
// 时间复杂度 O(n)，空间复杂度 O(k)

// 定义最大堆结构
type maxHeap struct {
	data []int // 存储元素索引的底层数组
	nums []int // 引用原始数组用于值比较
	k    int   // 滑动窗口大小
}

// 实现 heap.Interface 需要的方法
func (h *maxHeap) Len() int { return len(h.data) }

// 比较规则：实现最大堆（值大的排在前面）
func (h *maxHeap) Less(i, j int) bool {
	return h.nums[h.data[i]] > h.nums[h.data[j]]
}

// 交换元素位置
func (h *maxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 压入元素时的优化处理（核心逻辑）
func (h *maxHeap) Push(x any) {
	idx := x.(int) // 当前元素的索引
	// 关键优化：维护单调性 - 移除所有比当前元素小的旧元素
	// 因为这些旧元素不可能成为后续窗口的最大值
	for h.Len() > 0 && h.nums[idx] >= h.nums[h.data[len(h.data)-1]] {
		h.data = h.data[:len(h.data)-1] // 从尾部移除无效元素
	}
	h.data = append(h.data, idx) // 添加当前元素索引
}

// 弹出元素（常规实现）
func (h *maxHeap) Pop() any {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}

// 修剪过期元素（维护窗口有效性）
func (h *maxHeap) trim(current int) {
	// 计算窗口左边界：current - k + 1
	left := current - h.k + 1
	// 移除所有超出左边界的元素（这些元素已不在当前窗口内）
	for h.Len() > 0 && h.data[0] < left {
		h.data = h.data[1:] // 直接操作底层数组，比heap.Pop更高效
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	// 初始化堆，预分配容量为k+1（减少内存分配次数）
	h := &maxHeap{
		nums: nums,
		k:    k,
		data: make([]int, 0, k+1),
	}
	heap.Init(h) // 初始化堆结构

	result := make([]int, 0, len(nums)-k+1) // 预分配结果数组

	for i := 0; i < len(nums); i++ {
		// 1. 将当前索引压入堆（自动维护单调性）
		heap.Push(h, i)

		// 2. 修剪过期元素（保持堆顶始终在窗口内）
		h.trim(i)

		// 3. 当窗口形成后（i >= k-1），记录当前窗口最大值
		if i >= k-1 {
			result = append(result, nums[h.data[0]])
		}
	}
	return result
}

// 215: 数组中的第 K 个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
func findKthLargest(nums []int, k int) int {
	// 初始化堆的大小为数组长度
	heapSize := len(nums)
	// 构建最大堆
	buildMaxHeap(nums, heapSize)

	// 执行 k-1 次删除堆顶操作
	// 每次将堆顶（最大值）放到数组末尾，并减小堆的大小
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		// 交换堆顶元素和当前位置的元素
		nums[0], nums[i] = nums[i], nums[0]
		// 减小堆的大小
		heapSize--
		// 对新的堆顶元素进行下沉操作，维护最大堆性质
		maxHeapify(nums, 0, heapSize)
	}
	// 返回当前堆顶元素，即第k个最大元素
	return nums[0]
}

// buildMaxHeap 构建最大堆
// a: 要构建堆的数组
// heapSize: 堆的大小
func buildMaxHeap(a []int, heapSize int) {
	// 从最后一个非叶子节点开始，自底向上构建最大堆
	// 最后一个非叶子节点的索引是 heapSize/2 - 1
	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

// maxHeapify 维护最大堆的性质
// a: 堆数组
// i: 当前需要维护的节点索引
// heapSize: 堆的大小
func maxHeapify(a []int, i, heapSize int) {
	// 计算左子节点(2i+1)、右子节点(2i+2)的索引，初始化最大值索引为当前节点
	l, r, largest := i*2+1, i*2+2, i

	// 如果左子节点存在且大于当前最大值，更新最大值索引
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	// 如果右子节点存在且大于当前最大值，更新最大值索引
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}

	// 如果最大值不是当前节点，需要交换并继续维护堆的性质
	if largest != i {
		// 交换当前节点和最大值节点
		a[i], a[largest] = a[largest], a[i]
		// 递归维护被交换的子树
		maxHeapify(a, largest, heapSize)
	}
}

func main() {
	mergeKLists([]*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	})

	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) // 输出 5
}
