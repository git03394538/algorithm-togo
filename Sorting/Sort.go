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

func BubbleSort(input []int) []int {
	for k := 0; k < len(input) - 2; k++ {
		for i := 0; i < len(input) - 1; i++ {
			if input[i] > input[i + 1] {
				input[i], input[i + 1] = input[i + 1], input[i]
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
		input[k], input[smallestIndex] = input[smallestIndex], input[k]
	}
	return input
}

func InsertionSort(input []int) []int {
	for i := range input {
		preIndex := i - 1
		current := input[i]
		for preIndex >= 0 && input[preIndex] > current {
			input[preIndex+1] = input[preIndex]
			preIndex -= 1
		}
		input[preIndex+1] = current
	}
	return input
}


