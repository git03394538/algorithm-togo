package FibonacciPisano

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFibonacciPisano(t *testing.T) {
	assert.Equal(t, getFibonacciPisano(1, 239), 1)
	assert.Equal(t, getFibonacciPisano(2015, 3), 1)
	assert.Equal(t, getFibonacciPisano(239, 1000), 161)
	assert.Equal(t, getFibonacciPisano(2816213588, 30524), 10249)
}