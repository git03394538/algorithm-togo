package FibonacciPartialSum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciPartialSum(t *testing.T) {
	assert.Equal(t, getFibonacciPartialSum(3, 7), 1)
	assert.Equal(t, getFibonacciPartialSum(10, 10), 5)
	assert.Equal(t, getFibonacciPartialSum(10, 200), 2)
}

