package LeastCommonMultiple

import "Algorithm/Week1/Gcd"

func getLCMSlow(a int, b int) int {
	for l := 1; l <= a * b; l++ {
		if l % a == 0 && l % b == 0 {
			return l
		}
	}
	return a * b
}

func getLCMFast(a int, b int) int {
	return a * b / Gcd.GetGcdFast(a, b)
}


