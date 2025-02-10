package sort

import "learn-go/structure/constraints"

// Bucket 函数实现了桶排序算法，适用于数字类型的切片。
// T 是一个类型约束，表示该函数可以处理任何实现了 constraints.Number 接口的类型。
func Bucket[T constraints.Number](arr []T) []T {
	// 如果数组长度小于等于 1，直接返回原数组
	if len(arr) <= 1 {
		return arr
	}

	// 初始化最大值和最小值
	max := arr[0] // 假设第一个元素为最大值
	min := arr[0] // 假设第一个元素为最小值

	// 遍历数组，找到最大值和最小值
	for _, v := range arr {
		if v > max {
			max = v // 更新最大值
		}
		if v < min {
			min = v // 更新最小值
		}
	}

	// 创建桶，桶的数量与数组长度相同
	bucket := make([][]T, len(arr))

	// 将每个元素放入对应的桶中
	for _, v := range arr {
		// 计算桶的索引
		// 使用 (v - min) / (max - min) 计算相对位置，并映射到桶的索引
		bucketIndex := int((v - min) / (max - min) * T(len(arr)-1))
		bucket[bucketIndex] = append(bucket[bucketIndex], v) // 将元素添加到对应的桶中
	}

	// 对每个桶进行插入排序
	for i := range bucket {
		bucket[i] = Insertion(bucket[i]) // 假设 Insertion 是一个插入排序的实现
	}

	// 创建一个切片用于存储排序后的结果
	sorted := make([]T, 0, len(arr))

	// 将所有桶中的元素合并到 sorted 切片中
	for _, v := range bucket {
		sorted = append(sorted, v...) // 将桶中的元素添加到结果切片
	}

	return sorted // 返回排序后的切片
}
