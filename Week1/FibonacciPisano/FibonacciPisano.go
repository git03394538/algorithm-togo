package FibonacciPisano

func getFibonacciPisano(n int, m int) int {
	if n == 1 {
		return n
	}

	pisonoPeriod := -1
	pisonaArray := []int{0, 1, 1}

	for i := 3; i <= n; i ++ {
		x := pisonaArray[i - 1]
		y := pisonaArray[i - 2]
		pisonaArray = append(pisonaArray, (x + y) % m)
		if pisonaArray[i - 1] == 0 && pisonaArray[i] == 1 {
			pisonoPeriod = i - 1
			break
		}
	}
	if pisonoPeriod == -1 {
		return pisonaArray[n]
	}
	return pisonaArray[n % pisonoPeriod] % m
}
