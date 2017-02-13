package FibonacciLastDigit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetFibonacciLastDigit(t *testing.T) {
	assert.Equal(t, getFibonacciLastDigit(331), 9)
	assert.Equal(t, getFibonacciLastDigit(327305), 5)
}
