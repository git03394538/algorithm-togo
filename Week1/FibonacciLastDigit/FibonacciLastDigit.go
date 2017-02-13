package FibonacciLastDigit

func getFibonacciLastDigit(n int) int {
	if n <= 1 {
		return n
	}
	previous := 0
	current := 1

	for i := 0; i < n - 1; i++ {
		var tmp_previous = previous
		previous = current
		current = (tmp_previous + current) % 10
	}

	return current
}
