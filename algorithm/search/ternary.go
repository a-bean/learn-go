package search

import (
	"fmt"
	"math"
)

// TernaryMax 使用三分搜索算法在区间[a,b]上寻找单峰函数f(x)的最大值
// 参数:
//
//	a, b    - 搜索区间端点（需有限值）
//	epsilon - 精度阈值（区间长度终止条件）
//	f       - 目标函数（需在区间内单峰）
//
// 返回值:
//
//	float64 - 找到的最大值
//	error   - 错误信息（当区间端点为无穷大时返回错误）
//
// 原理:
//  1. 将区间分为三等分：left=(2a+b)/3, right=(a+2b)/3
//  2. 比较f(left)和f(right)的值
//  3. 保留包含更大值的区间（若f(left)<f(right)则保留右区间）
//  4. 递归直到区间长度小于epsilon
func TernaryMax(a, b, epsilon float64, f func(x float64) float64) (float64, error) {
	if a == math.Inf(-1) || b == math.Inf(1) {
		return -1, fmt.Errorf("interval boundaries should be finite numbers")
	}

	// 终止条件：区间长度达到精度要求
	if math.Abs(a-b) <= epsilon {
		return f((a + b) / 2), nil
	}

	// 计算两个三分点
	left := (2*a + b) / 3
	right := (a + 2*b) / 3

	// 根据函数值决定保留区间
	if f(left) < f(right) {
		// 右区间包含更大值，保留[right, b]
		return TernaryMax(left, b, epsilon, f)
	}
	// 左区间包含更大值，保留[a, right]
	return TernaryMax(a, right, epsilon, f)
}

// TernaryMin 使用三分搜索算法在区间[a,b]上寻找单峰函数f(x)的最小值
// 参数和返回值同TernaryMax，但寻找方向相反
func TernaryMin(a, b, epsilon float64, f func(x float64) float64) (float64, error) {
	if a == math.Inf(-1) || b == math.Inf(1) {
		return -1, fmt.Errorf("interval boundaries should be finite numbers")
	}

	if math.Abs(a-b) <= epsilon {
		return f((a + b) / 2), nil
	}

	left := (2*a + b) / 3
	right := (a + 2*b) / 3

	// 比较方向与TernaryMax相反
	if f(left) > f(right) {
		// 右区间包含更小值，保留[right, b]
		return TernaryMin(left, b, epsilon, f)
	}
	// 左区间包含更小值，保留[a, right]
	return TernaryMin(a, right, epsilon, f)
}
