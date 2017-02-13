package Fibonacci

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciReturnTheSame(t *testing.T) {
	assert.Equal(t, getFibonacciSlow(20), getFibonacciFast(20))
}
