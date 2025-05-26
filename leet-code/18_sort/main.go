package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
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
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 493 翻转对 https://leetcode.cn/problems/reverse-pairs/
// reversePairs 函数计算数组中逆序对的数量
func reversePairs(nums []int) int {
	buf := make([]int, len(nums))    // 创建一个缓冲数组用于合并
	return mergeSortCount(nums, buf) // 调用归并排序并返回逆序对的数量
}
func mergeSortCount(nums, buf []int) int {
	if len(nums) <= 1 { // 如果数组长度小于等于1，返回0，因为没有逆序对
		return 0
	}
	mid := (len(nums) - 1) / 2               // 计算中间索引
	cnt := mergeSortCount(nums[:mid+1], buf) // 递归计算左半部分的逆序对数量
	cnt += mergeSortCount(nums[mid+1:], buf) // 递归计算右半部分的逆序对数量

	// 计算逆序对
	for i, j := 0, mid+1; i < mid+1; i++ { // 遍历左半部分
		// Note!!! j 是递增的
		for ; j < len(nums) && nums[i] <= 2*nums[j]; j++ { // 找到 nums[i] > 2 * nums[j] 的所有 j
		}
		cnt += len(nums) - j // 统计逆序对数量
	}

	copy(buf, nums) // 将当前数组复制到缓冲数组中
	// 合并两个已排序的子数组
	for i, j, k := 0, mid+1, 0; k < len(nums); {
		if j >= len(nums) || i < mid+1 && buf[i] > buf[j] { // 如果左侧元素大于右侧元素
			nums[k] = buf[i] // 将左侧元素放入原数组
			i++              // 移动左侧指针
		} else {
			nums[k] = buf[j] // 将右侧元素放入原数组
			j++              // 移动右侧指针
		}
		k++ // 移动合并数组指针
	}
	return cnt // 返回逆序对的数量
}

// 327 区间和的个数 https://leetcode.cn/problems/count-of-range-sum/

// 147 对链表进行插入排序 https://leetcode.cn/problems/insertion-sort-list/
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: nil} // 这里初始化不要直接指向 head，为了下面循环可以统一处理
	cur, pre := head, newHead
	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		cur = next

		pre = newHead // 归位，重头开始
	}
	return newHead.Next
}

// 148 排序链表 https://leetcode.cn/problems/sort-list/
// merge1 函数合并两个已排序的链表并返回合并后的链表
func merge1(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}                      // 创建一个虚拟头节点
	temp, temp1, temp2 := dummyHead, head1, head2 // 初始化指针

	// 遍历两个链表，合并它们
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val { // 比较当前节点的值
			temp.Next = temp1  // 将较小的节点连接到合并链表
			temp1 = temp1.Next // 移动到下一个节点
		} else {
			temp.Next = temp2  // 将较小的节点连接到合并链表
			temp2 = temp2.Next // 移动到下一个节点
		}
		temp = temp.Next // 移动合并链表的指针
	}

	// 处理剩余节点
	if temp1 != nil {
		temp.Next = temp1 // 如果链表1还有剩余节点，直接连接
	} else if temp2 != nil {
		temp.Next = temp2 // 如果链表2还有剩余节点，直接连接
	}
	return dummyHead.Next // 返回合并后的链表，跳过虚拟头节点
}

// sort1 函数使用归并排序算法对链表进行排序
func sort1(head, tail *ListNode) *ListNode {
	if head == nil { // 如果链表为空，返回空
		return head
	}

	if head.Next == tail { // 如果链表只有一个节点，返回该节点
		head.Next = nil
		return head
	}

	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next // 慢指针每次移动一步
		fast = fast.Next // 快指针每次移动两步
		if fast != tail {
			fast = fast.Next // 确保快指针不越界
		}
	}

	mid := slow // 中间节点
	// 递归排序左右两部分并合并
	return merge1(sort1(head, mid), sort1(mid, tail))
}

// sortList 函数是对外接口，调用 sort1 进行排序
func sortList(head *ListNode) *ListNode {
	return sort1(head, nil) // 从头节点开始排序
}

// 164 最大间距 https://leetcode.cn/problems/maximum-gap/
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0 // 如果元素少于2，返回0
	}

	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		if num < minVal {
			minVal = num // 找到最小值
		}
		if num > maxVal {
			maxVal = num // 找到最大值
		}
	}

	// 计算桶的大小
	bucketSize := max(1, (maxVal-minVal)/(len(nums)-1)) // 每个桶的大小
	bucketCount := (maxVal-minVal)/bucketSize + 1       // 桶的数量

	// 创建桶
	buckets := make([][2]int, bucketCount) // 每个桶保存[min, max]
	for i := range buckets {
		buckets[i][0] = math.MaxInt32 // 初始化最小值为最大整数
		buckets[i][1] = math.MinInt32 // 初始化最大值为最小整数
	}

	// 将数字分配到桶中
	for _, num := range nums {
		idx := (num - minVal) / bucketSize // 确定桶的索引
		if num < buckets[idx][0] {
			buckets[idx][0] = num // 更新桶的最小值
		}
		if num > buckets[idx][1] {
			buckets[idx][1] = num // 更新桶的最大值
		}
	}

	// 计算最大间隔
	maxGap := 0
	previousMax := buckets[0][1] // 从第一个桶的最大值开始
	for i := 1; i < bucketCount; i++ {
		if buckets[i][0] == math.MaxInt32 { // 跳过空桶
			continue
		}
		maxGap = max(maxGap, buckets[i][0]-previousMax) // 更新最大间隔
		previousMax = buckets[i][1]                     // 更新前一个桶的最大值
	}

	return maxGap // 返回找到的最大间隔
}

// 179 最大数 https://leetcode.cn/problems/largest-number/
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

// 220 存在重复元素 III https://leetcode.cn/problems/contains-duplicate-iii/
// containsNearbyAlmostDuplicate 函数检查数组中是否存在两个不同的索引 i 和 j，使得
// |nums[i] - nums[j]| <= t 且 |i - j| <= k。
// 该函数使用桶排序的思想来实现高效查找。
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	// 检查边界条件
	if k <= 0 || t < 0 || len(nums) < 2 {
		return false // 如果 k <= 0 或 t < 0 或数组长度小于2，返回 false
	}

	buckets := map[int]int{} // 创建一个桶，用于存储元素及其值

	for i := 0; i < len(nums); i++ {
		// 计算当前元素的桶索引
		key := nums[i] / (t + 1) // 使用 (t + 1) 来避免桶重叠
		if nums[i] < 0 {
			key-- // 如果元素为负数，调整桶索引
		}

		// 检查当前桶是否已存在
		if _, ok := buckets[key]; ok {
			return true // 如果当前桶已存在，说明找到了满足条件的元素
		}

		// 检查左侧桶
		if v, ok := buckets[key-1]; ok && nums[i]-v <= t {
			return true // 如果左侧桶存在且满足条件，返回 true
		}

		// 检查右侧桶
		if v, ok := buckets[key+1]; ok && v-nums[i] <= t {
			return true // 如果右侧桶存在且满足条件，返回 true
		}

		// 如果桶的数量超过 k，删除最旧的桶
		if len(buckets) >= k {
			delete(buckets, nums[i-k]/(t+1)) // 删除超出范围的桶
		}

		// 将当前元素放入桶中
		buckets[key] = nums[i]
	}
	return false // 如果没有找到满足条件的元素，返回 false
}

// 189. 轮转数组 https://leetcode.cn/problems/rotate-array/
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n             // 处理 k 大于 n 的情况
	reverse(nums, 0, n-1) // 反转整个数组
	reverse(nums, 0, k-1) // 反转前 k 个元素
	reverse(nums, k, n-1) // 反转后 n-k 个元素
}
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start] // 交换元素
		start++                                         // 移动指针
		end--                                           // 移动指针
	}
}

// 324 摆动排序 II https://leetcode.cn/problems/wiggle-sort-ii/
// 767 重构字符串 https://leetcode.cn/problems/reorganize-string/
// 969 灯泡开关 https://leetcode.cn/problems/pancake-sorting/
// 1054 距离相等的数组对 https://leetcode.cn/problems/distinct-echo-substrings/

// 215 数组中的第 K 个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/
// findKthLargest 查找数组中第k个最大的元素
// nums: 输入数组
// k: 要找的第k个最大元素
// 返回：第k个最大元素的值
// 时间复杂度：平均O(n)，最坏O(n²)
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	// 将第k大转换为第(n-k)小，这样可以统一处理
	return quickselect(nums, 0, n-1, n-k)
}

// quickselect 快速选择算法的核心实现
// nums: 待处理的数组
// l: 左边界
// r: 右边界
// k: 要找的下标（第k小的数）
func quickselect(nums []int, l, r, k int) int {
	// 如果区间只有一个元素，且是我们要找的k，直接返回
	if l == r {
		return nums[k]
	}

	// 选择第一个元素作为分区点(pivot)
	partition := nums[l]
	// i指针从左边开始，j指针从右边开始
	i := l - 1
	j := r + 1

	// 双指针分区过程
	for i < j {
		// 找到第一个大于等于partition的元素
		for i++; nums[i] < partition; i++ {
		}
		// 找到第一个小于等于partition的元素
		for j--; nums[j] > partition; j-- {
		}
		// 如果i和j没有相遇，交换这两个元素
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 根据分区点位置决定向左还是向右继续查找
	if k <= j {
		// 如果k在分区点左边（包含分区点），继续在左半部分查找
		return quickselect(nums, l, j, k)
	} else {
		// 如果k在分区点右边，继续在右半部分查找
		return quickselect(nums, j+1, r, k)
	}
}

func main() {
	relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	relativeSortArray1([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6})
	merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	reversePairs([]int{7, 5, 6, 4})

	insertionSortList(&ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}})

	sortList(&ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}})

	maximumGap([]int{3, 6, 9, 1})

	largestNumber([]int{3, 30, 34, 5, 9})

	containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0)
	fmt.Print(findKthLargest([]int{3, 2, 1, 5, 6, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 2))
}
