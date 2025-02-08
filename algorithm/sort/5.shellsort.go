package sort

import (
	"learn-go/structure/constraints"
)

// 插入排序优化: 希尔排序
func Shell[T constraints.Ordered](arr []T) []T {
	for d := int(len(arr) / 2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	return arr
}
