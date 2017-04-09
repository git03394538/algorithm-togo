package DPChange

func RecursiveChange(amount int, coins []int) int {
	if amount == 0 {
		return 0
	}

	best := -1
	var nextTry int
	for _,coin := range coins {
		if coin <= amount {
			nextTry = RecursiveChange(amount - coin, coins)
		}

		if best < 0 || best > nextTry + 1 {
			best = nextTry + 1
		}

	}
	return best
}


func DPChange(amount int, coins []int) int {
	results := []int{0}
	var best int
	for m := 1; m <= amount; m++ {
		results = append(results, m)
		for _,coin := range coins {
			if coin <= m {
				best = results[m - coin] + 1
			}

			if best < results[m] {
				results[m] = best
			}

		}
	}
	return results[amount]
}
