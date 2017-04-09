package DP2

func EditDistance(s1 string, s2 string) int {
	len1 := len([]rune(s1))
	len2 := len([]rune(s2))


	results := make([][]int, len1 + 1)
	for k := range results {
		results[k] = make([]int, len2 + 1)
	}

	for i := 0; i <= len1; i++ {
		results[i][0] = i
	}

	for j := 0; j <= len2; j++ {
		results[0][j] = j
	}

	for i := 0; i < len1; i++ {
		c1 := []rune(s1)[i]
		for j:= 0; j < len2; j++ {
			c2 := []rune(s2)[j]

			if c1 == c2 {
				results[i + 1][j + 1] = results[i][j]
			} else {
				replaced := results[i][j] + 1
				inserted := results[i][j + 1] + 1
				deleted := results[i + 1][j] + 1

				var min int
				if replaced > inserted {
					min = inserted
				} else {
					min = replaced
				}
				if !(deleted > min) {
					min = deleted
				}
				results[i + 1][j + 1] = min
			}
		}
	}

	return results[len1][len2]
}