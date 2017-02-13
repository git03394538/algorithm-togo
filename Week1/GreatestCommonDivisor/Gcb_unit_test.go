package GreatestCommonDivisor

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlowGcb(t *testing.T) {
	assert.Equal(t, getGcdSlow(10, 4), 2)
}

func TestFastGcd(t *testing.T) {
	assert.Equal(t, GetGcdFast(357, 234), 3)
}

func TestFastAndSlowIsTheSame(t *testing.T) {
	assert.Equal(t, GetGcdFast(x, y), getGcdSlow(x, y))
}
