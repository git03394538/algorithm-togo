package FibonacciLastDigitSum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciLastDigitSum(t *testing.T) {
	assert.Equal(t, getFibonacciLastDigitSum(3), 4)
	assert.Equal(t, getFibonacciLastDigitSum(100), 5)
	assert.Equal(t, getFibonacciLastDigitSum(99999999999999999), 0)
}
