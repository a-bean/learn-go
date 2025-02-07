package search

// 插值搜索算法是一种改进的二分搜索算法，主要用于在 均匀分布 的 有序 数组中快速查找目标值。它的核心思想是通过目标值与数组中值的相对位置来预测目标值可能的位置，而不是像二分搜索那样简单地取中间位置。

// 算法原理：
// 1. 基本思想：假设数组中的值是均匀分布的，通过目标值与当前搜索范围的最小值和最大值的比例，预测目标值可能的位置。

// 2. 公式： mid = low + ((guess-sortedData[low]) * (high-low)) / (sortedData[high] - sortedData[low])

// 3. 搜索过程：
// 如果 sortedData[mid] == guess，则找到目标值。
// 如果 sortedData[mid] > guess，说明目标值在左半部分，调整上界 high = mid - 1。
// 如果 sortedData[mid] < guess，说明目标值在右半部分，调整下界 low = mid + 1。

// 4.终止条件：
// 找到目标值。
// 搜索范围缩小到无效（low > high），表示目标值不存在。

// 5.适用场景：
// 数据量大且分布均匀的有序数组。
// 目标值的范围已知且分布均匀。

// 插值搜索算法
func Interpolation(sortedData []int, guess int) (int, error) {
	// 如果数组为空，直接返回错误
	if len(sortedData) == 0 {
		return -1, ErrNotFound
	}

	// 初始化搜索范围的上下界及对应的值
	var (
		low, high       = 0, len(sortedData) - 1
		lowVal, highVal = sortedData[low], sortedData[high]
	)

	// 当上下界值不同且猜测值在范围内时进行插值搜索
	for lowVal != highVal && (lowVal <= guess) && (guess <= highVal) {
		// 计算插值位置
		mid := low + int(float64(float64((guess-lowVal)*(high-low))/float64(highVal-lowVal)))

		// 如果找到目标值，返回第一个匹配的位置
		if sortedData[mid] == guess {
			for mid > 0 && sortedData[mid-1] == guess {
				mid--
			}
			return mid, nil
		}

		// 如果中间值大于猜测值，调整上界
		if sortedData[mid] > guess {
			high, highVal = mid-1, sortedData[high]
		} else {
			// 否则调整下界
			low, lowVal = mid+1, sortedData[low]
		}
	}

	// 检查下界是否等于猜测值
	if guess == lowVal {
		return low, nil
	}
	// 未找到目标值
	return -1, ErrNotFound
}
