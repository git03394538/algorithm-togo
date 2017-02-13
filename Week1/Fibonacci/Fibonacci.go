package Fibonacci

func getFibonacciSlow(n int) int {
	if n <= 1 {
		return n
	} else {
		return getFibonacciSlow(n - 1) + getFibonacciSlow(n - 2)
	}
}

func getFibonacciFast(n int) int {
	f := []int{}
	f = append(f, 0)
	f = append(f, 1)

	for i := 2; i <= n; i++ {
		f = append(f, f[i - 1] + f [i - 2])
	}
	return f[n]
}