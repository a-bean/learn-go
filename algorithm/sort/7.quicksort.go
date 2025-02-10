package sort

import (
	"learn-go/structure/constraints"
)

// Partition 函数实现了快速排序中的分区操作。
// 它将数组 arr 中的元素根据 pivotElement 进行分区，
// 并返回新的 pivot 索引，使得左侧的元素都小于等于 pivot，右侧的元素都大于 pivot。
func Partition[T constraints.Ordered](arr []T, low, high int) int {
	index := low - 1
	pivotElement := arr[high]

	// 遍历数组，将小于等于 pivot 的元素移动到左侧
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}

	// 将 pivot 元素放到正确的位置
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1 // 返回 pivot 的新索引
}

func Partition1[T constraints.Ordered](arr []T, low, high int) int {
	// 使用三数取中法选择 pivot
	mid := low + (high-low)/2
	if arr[low] > arr[mid] {
		arr[low], arr[mid] = arr[mid], arr[low]
	}
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	if arr[mid] > arr[high] {
		arr[mid], arr[high] = arr[high], arr[mid]
	}

	// 将 pivot 放到最后
	pivotIndex := mid
	pivotValue := arr[pivotIndex]
	arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

	// 双指针法进行分区
	storeIndex := low
	for i := low; i < high; i++ {
		if arr[i] < pivotValue {
			arr[storeIndex], arr[i] = arr[i], arr[storeIndex]
			storeIndex++
		}
	}

	// 将 pivot 放到正确的位置
	arr[storeIndex], arr[high] = arr[high], arr[storeIndex]
	return storeIndex // 返回 pivot 的新索引
}

// QuicksortRange 函数实现了快速排序的递归逻辑。
// 它对数组 arr 的指定范围 [low, high] 进行排序。
func QuicksortRange[T constraints.Ordered](arr []T, low, high int) {
	if len(arr) <= 1 { // 如果数组长度小于等于 1，直接返回
		return
	}

	if low < high { // 确保 low 小于 high
		pivot := Partition(arr, low, high) // 进行分区操作，获取 pivot 索引
		// 递归对左侧和右侧的子数组进行排序
		QuicksortRange(arr, low, pivot-1)  // 排序左侧子数组
		QuicksortRange(arr, pivot+1, high) // 排序右侧子数组
	}
}

// Quicksort 函数是快速排序的入口函数。
// 它对整个数组 arr 进行排序，并返回排序后的数组。
func Quicksort[T constraints.Ordered](arr []T) []T {
	QuicksortRange(arr, 0, len(arr)-1) // 调用 QuicksortRange 进行排序
	return arr                         // 返回排序后的数组
}
