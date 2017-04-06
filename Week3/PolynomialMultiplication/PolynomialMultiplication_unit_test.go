package PolynomialMultiplication_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"Algorithm/Week3/PolynomialMultiplication"
)

func TestPolynomialMultiplication(t *testing.T) {
	assert.Equal(t, []int{15, 13, 33, 9, 10}, PolynomialMultiplication.Calculate([]int{3, 2, 5}, []int{5, 1, 2}, 3))
	assert.Equal(t, []int{4, 11, 20, 30, 20, 11, 4}, PolynomialMultiplication.Calculate([]int{4, 3, 2, 1}, []int{1, 2, 3, 4}, 4))
}