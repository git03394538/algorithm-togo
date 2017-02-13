package GreatestCommonDivisor

func getGcdSlow(a int, b int) int {
	best := 0
	for d := 1; d <= a + b; d++ {
		if a % d == 0 && b % d == 0 {
			best = d
		}
	}
	return best
}

func GetGcdFast(a int, b int) int {
	if b == 0 {
		return a
	}
	return GetGcdFast(b, a % b)
}