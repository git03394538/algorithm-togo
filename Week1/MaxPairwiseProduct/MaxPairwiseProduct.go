package MaxPairwiseProduct

func getMaxPairwiseProductSlow(input []int) int64 {
	result := int64(0)
	n := len(input)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if int64(input[i]) * int64(input[j]) > result {
				result = int64(input[i]) * int64(input[j])
			}
		}
	}
	return result
}

func getMaxPairwiseProductNormal(input []int) int64 {
	n := len(input)
	index1 := -1
	index2 := -1

	for i := 0; i < n; i++ {
		if index1 == -1 || input[i] > input[index1] {
			index1 = i
		}
	}

	for j := 0; j < n; j++ {
		if index1 != j && (index2 == -1 || input[j] > input[index2]) {
			index2 = j
		}
	}
	return int64(input[index1]) * int64(input[index2])
}

func getMaxPairwiseProductFast(input []int) int64 {
	n := len(input)
	index1 := 0
	index2 := 1

	if input[index1] < input[index2] {
		index2 = 0
		index1 = 1
	}
	for i := 2; i < n; i++ {
		if input[i] > input[index1] {
			index2 = index1
			index1 = i
		} else {
			if input[i] > input[index2] {
				index2 = i
			}
		}
	}
	return int64(input[index1]) * int64(input[index2])
}
