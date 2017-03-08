package Sorting

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	return combineArray(
		MergeSort(input[: len(input)/2]),
		MergeSort(input[len(input)/2 :]),
		len(input))
}

func combineArray(a []int, b []int, size int) []int {
	result := make([]int, size)
	ai := 0
	bi := 0

	for k := 0; k < size; k++ {
		if bi >= len(b) {
			result[k] = a[ai]
			ai++
		} else if ai >= len(a) || b[bi] < a[ai]{
			result[k] = b[bi]
			bi++
		} else {
			result[k] = a[ai]
			ai++
		}
	}
	return result
}

func MergeSortFloat64(input []float64) []float64 {
	if len(input) < 2 {
		return input
	}
	return combineArrayFloat64(
		MergeSortFloat64(input[: len(input)/2]),
		MergeSortFloat64(input[len(input)/2 :]),
		len(input))
}

func combineArrayFloat64(a []float64, b []float64, size int) []float64 {
	result := make([]float64, size)
	ai := 0
	bi := 0

	for k := 0; k < size; k++ {
		if bi >= len(b) {
			result[k] = a[ai]
			ai++
		} else if ai >= len(a) || b[bi] < a[ai]{
			result[k] = b[bi]
			bi++
		} else {
			result[k] = a[ai]
			ai++
		}
	}
	return result
}

func BubbleSort(input []int) []int {
	var temp int
	for k := 0; k < len(input) - 2; k++ {
		for i := 0; i < len(input) - 1; i++ {
			if input[i] > input[i + 1] {
				temp = input[i]
				input[i] = input[i + 1]
				input[i + 1] = temp
			}
		}
	}
	return input
}

func SelectionSort(input []int) []int {
	for k := 0; k < len(input); k++ {
		smallestIndex := k
		for i := k; i < len(input); i++ {
			if input[smallestIndex] > input[i] {
				smallestIndex = i
			}
		}
		temp := input[k]
		input[k] = input[smallestIndex]
		input[smallestIndex] = temp
	}
	return input
}


