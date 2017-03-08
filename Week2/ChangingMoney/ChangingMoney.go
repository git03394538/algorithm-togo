package ChangingMoney

func change(amount int) int {
	coins := []int{10, 5, 1}
	result := 0
	for _,coin := range coins {
		for amount - coin >= 0 {
			result += 1
			amount -= coin
		}
	}

	return result
}
