// 跳跃搜索算法原理
// 跳跃搜索（Jump Search）是一种用于有序数组的搜索算法，它通过跳跃式地检查数组中的元素来缩小搜索范围，然后在较小的范围内进行线性搜索。它的时间复杂度介于线性搜索和二分搜索之间

// 算法步骤：
// 1. 确定跳跃步长：
// 跳跃步长通常为数组长度的平方根，即 step= 根号n 。
// 这是为了在跳跃次数和线性搜索范围之间取得平衡。
// 2. 跳跃阶段：
// 从数组的起始位置开始，每次跳跃 step 个元素，检查当前元素是否大于目标值。
// 如果当前元素大于目标值，说明目标值可能在前一个跳跃点和当前跳跃点之间。
// 如果跳跃超出数组范围，则将当前跳跃点调整为数组的最后一个元素。
// 3. 线性搜索阶段：
// 在确定的范围内（即前一个跳跃点到当前跳跃点之间）进行线性搜索，逐个检查元素是否等于目标值。
// 如果找到目标值，返回其索引；否则返回未找到。

// 适用场景：
// 数据量中等且有序的数组。
// 需要简单实现的搜索场景。
// 不适合数据量非常大或无序的情况。

package search

import "math"

// Jump 搜索通过跳跃多个步骤在有序列表中前进，直到找到大于目标值的项，
// 然后从最后搜索的项到当前项创建一个子列表并执行线性搜索。
func Jump(array []int, target int) (int, error) {
	n := len(array)
	if n == 0 {
		return -1, ErrNotFound
	}

	// 步长的最优值是列表长度的平方根
	step := int(math.Round(math.Sqrt(float64(n))))

	prev := 0    // 前一个索引
	curr := step // 当前索引

	for array[curr-1] < target {
		prev = curr
		if prev >= len(array) {
			return -1, ErrNotFound
		}

		curr += step

		// 防止跳跃超过列表范围
		if curr > n {
			curr = n
		}
	}

	// 从索引 prev 到索引 curr 执行线性搜索
	for array[prev] < target {
		prev++

		// 如果到达范围末尾，表示未找到目标值
		if prev == curr {
			return -1, ErrNotFound
		}
	}
	if array[prev] == target {
		return prev, nil
	}

	return -1, ErrNotFound
}
