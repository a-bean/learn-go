package search

// Binary 是递归实现，适合理解二分查找的基本原理。
func Binary(array []int, target int, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}

// BinaryIterative 是迭代实现，适合实际应用。
func BinaryIterative(array []int, target int) (int, error) {
	startIndex := 0
	endIndex := len(array) - 1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if array[mid] > target {
			endIndex = mid - 1
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}

// LowerBound 返回数组中第一个不小于目标值的元素索引
// 使用二分查找算法在有序数组中查找
// 如果找到符合条件的元素，返回其索引；否则返回-1和ErrNotFound
func LowerBound(array []int, target int) (int, error) {
	n := len(array)
	if n == 0 {
		return -1, ErrNotFound
	}

	start, end := 0, n // 使用左闭右开区间[0, n)
	for start < end {  // 终止条件更明确
		mid := start + (end-start)/2 // 避免溢出
		if array[mid] < target {
			start = mid + 1 // 收缩左边界
		} else {
			end = mid // 收缩右边界
		}
	}

	// 最终start指向第一个>=target的位置
	if start >= n || array[start] < target {
		return -1, ErrNotFound
	}
	return start, nil
}

// UpperBound 返回数组中第一个大于目标值的元素索引
// 使用二分查找算法在有序数组中查找
// 如果找到符合条件的元素，返回其索引；否则返回-1和ErrNotFound
func UpperBound(array []int, target int) (int, error) {
	n := len(array)
	if n == 0 {
		return -1, ErrNotFound
	}

	start, end := 0, n // 使用左闭右开区间[0, n)
	for start < end {  // 终止条件更明确
		mid := start + (end-start)/2 // 避免溢出
		if array[mid] > target {
			end = mid // 收缩右边界
		} else {
			start = mid + 1 // 收缩左边界
		}
	}

	// 最终start指向第一个>target的位置
	if start == 0 || array[start-1] > target {
		return -1, ErrNotFound
	}
	return start - 1, nil // 返回最后一个<=target的位置
}
