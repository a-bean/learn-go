package sort

import (
	"learn-go/structure/constraints"
)

func maxInt[T constraints.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}

func Pigeonhole[T constraints.Integer](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}

	max := maxInt(arr...)
	min := minInt(arr...)

	size := max - min + 1

	holes := make([]T, size)

	for _, element := range arr {
		holes[element-min]++
	}

	i := 0

	for j := T(0); j < size; j++ {
		for holes[j] > 0 {
			holes[j]--
			arr[i] = j + min
			i++
		}
	}

	return arr
}
