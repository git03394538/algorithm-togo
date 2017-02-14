package FibonacciLastDigitSum

func getFibonacciLastDigitSum(n int) int {
	if n <= 1 {
		return n
	}

	pisonoPeriod := 60
	if n < 60 {
		pisonoPeriod = n
	} else {
		n = n % pisonoPeriod
	}

	pisonaArray := []int{0, 1, 1}
	pisonaSumArray := []int{0, 1, 1}
	sum := 2

	for i := 3; i <= 60; i ++ {
		x := pisonaArray[i - 1]
		y := pisonaArray[i - 2]
		pisonaArray = append(pisonaArray, (x + y) % 10)
		sum = (sum + pisonaArray[i]) % 10
		pisonaSumArray = append(pisonaSumArray, sum)
	}

	return pisonaSumArray[n]
}
