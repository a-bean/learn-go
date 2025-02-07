package search

// 算法原理说明（快速选择算法）
// 核心思想：基于快速排序的分区思想，但每次只处理包含目标的那一半数据
// 时间复杂度：
// 平均情况：O(n)
// 最坏情况：O(n²)（可通过随机化基准优化到平均O(n)）
// 工作流程：
// 1. 选择基准元素进行分区
// 2. 根据基准位置与目标位置的关系，选择继续处理左/右分区
// 3. 递归直到基准位置等于目标位置

// SelectK 基于快速选择算法实现查找数组中第K大的元素
// 参数:
//
//	array - 输入数组（会被修改）
//	k     - 第k大的位置（1-based）
//
// 返回值:
//
//	int  - 找到的元素值
//	error - 当k超过数组长度时返回错误
//
// 原理:
//  1. 将第k大转换为第(n-k)小的索引位置（0-based）
//  2. 使用快速选择算法在O(n)平均时间复杂度内找到目标元素
//  3. 通过Lomuto分区方案进行原地分区操作
func SelectK(array []int, k int) (int, error) {
	if k > len(array) {
		return -1, ErrNotFound
	}
	// 转换逻辑示例：
	// 数组长度n=5，找第2大的元素 => 等价于找第5-2=3小的元素（索引3）
	return selectK(array, 0, len(array), len(array)-k), nil
}

// selectK 递归执行快速选择的核心函数
// 参数:
//
//	l    - 当前处理区间左边界（包含）
//	r    - 当前处理区间右边界（不包含）
//	idx  - 目标元素的最终排序位置
//
// 实现细节:
//
//	每次分区后根据基准位置与目标位置的关系，仅处理相关分区
//	平均时间复杂度: O(n) 最坏情况: O(n²)
func selectK(array []int, l, r, idx int) int {
	// 执行分区操作，得到基准元素的最终位置
	index := partition(array, l, r)

	switch {
	case index == idx: // 基准位置正好是目标位置
		return array[index]
	case index < idx: // 目标在右分区（调整左边界）
		return selectK(array, index+1, r, idx)
	default: // 目标在左分区（调整右边界）
		return selectK(array, l, index, idx)
	}
}

// partition Lomuto分区方案实现
// 参数:
//
//	l - 分区区间左边界（包含）
//	r - 分区区间右边界（不包含）
//
// 返回值:
//
//	基准元素的最终位置
//
// 分区过程:
//  1. 选择第一个元素作为基准（可优化为随机选择）
//  2. 维护j指针指向第一个大于基准的元素位置
//  3. 遍历过程中将小于等于基准的元素交换到j左侧
//  4. 最后将基准交换到正确位置
func partition(array []int, l, r int) int {
	elem, j := array[l], l+1 // 基准选择与指针初始化

	// 遍历区间 [l+1, r) 的元素
	for i := l + 1; i < r; i++ {
		if array[i] <= elem {
			// 将小于等于基准的元素交换到j位置，并移动j指针
			array[i], array[j] = array[j], array[i]
			j++
		}
	}

	// 将基准元素交换到最终位置（j-1）
	array[l], array[j-1] = array[j-1], array[l]
	return j - 1 // 返回基准的最终位置
}
