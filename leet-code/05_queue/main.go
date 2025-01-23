package main

import "fmt"

// 239: 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/description/
// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0, k)
	ans := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		// 删除出界的选项
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// 删除比当前值小的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		// 插入新选项并维护单调性
		deque = append(deque, i)

		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}
	return ans
}

// 862: 和至少为 K 的最短子数组 https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/
// 返回 A 的最短的非空连续子数组的长度，该子数组的和 至少 为 K 。如果没有和至少为 K 的非空子数组，返回 -1
func shortestSubarray(A []int, K int) int {
	res, prefixSum := len(A)+1, make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		prefixSum[i+1] = prefixSum[i] + A[i]
	}
	// deque 中保存递增的 prefixSum 下标
	deque := []int{}
	for i := range prefixSum {
		// 下面这个循环希望能找到 [deque[0], i] 区间内累加和 >= K，如果找到了就更新答案
		for len(deque) > 0 && prefixSum[i]-prefixSum[deque[0]] >= K {
			length := i - deque[0]
			if res > length {
				res = length
			}
			// 找到第一个 deque[0] 能满足条件以后，就移除它，因为它是最短长度的子序列了
			deque = deque[1:]
		}
		// 下面这个循环希望能保证 prefixSum[deque[i]] 递增
		for len(deque) > 0 && prefixSum[i] <= prefixSum[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	if res <= len(A) {
		return res
	}
	return -1
}

// 622: 设计循环队列 https://leetcode.cn/problems/design-circular-queue/description/
type MyCircularQueue struct {
	size  int
	cap   int
	queue []int
	left  int
	right int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size:  0,
		cap:   k,
		queue: make([]int, k),
		left:  0,
		right: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.size++
	this.queue[this.right] = value
	this.right++
	this.right %= this.cap // 关键操作：确保right指针循环
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	this.left++
	this.left %= this.cap
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.left]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.right == 0 {
		return this.queue[this.cap-1]
	}
	return this.queue[this.right-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}
func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}
func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(shortestSubarray([]int{1, 3, -1, -3, 5, 3, 6, 7}, 2))

}
