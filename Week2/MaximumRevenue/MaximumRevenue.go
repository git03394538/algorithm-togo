package MaximumRevenue

import (
	"Algorithm/Sorting"
)

func getMaximumRevenue(size int, a []int, b []int) int {
	a = Sorting.MergeSort(a)
	b = Sorting.MergeSort(b)
	result := 0
	for i := size - 1; i >= 0; i-- {
		result += a[i] * b[i]
	}

	return result
}
