package sort

import "learn-go/structure/constraints"

func Cycle[T constraints.Number](arr []T) []T {
	counter, cycle, len := 0, 0, len(arr)
	if len <= 1 {
		return arr
	}

	for cycle = 0; cycle < len-1; cycle++ {
		elem := arr[cycle]
		pos := cycle
		for counter = cycle + 1; counter < len; counter++ {
			if arr[counter] < elem {
				pos++
			}
		}
		if pos == cycle {
			continue
		}
		for elem == arr[pos] {
			pos++
		}
		arr[pos], elem = elem, arr[pos]

		for pos != cycle {
			pos = cycle
			for counter = cycle + 1; counter < len; counter++ {
				if arr[counter] < elem {
					pos++
				}
			}
			for elem == arr[pos] {
				pos++
			}

			if elem != arr[pos] {
				arr[pos], elem = elem, arr[pos]
			}
		}
	}

	return arr
}
