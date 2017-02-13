package Fibonacci

import (
	"testing"
)

func BenchmarkSlowFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getFibonacciSlow(40)
	}
}

func BenchmarkNormalFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getFibonacciFast(40)
	}
}