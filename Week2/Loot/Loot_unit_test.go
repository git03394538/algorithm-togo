package Loot

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMaximumLoot(t *testing.T) {
	assert.Equal(t, getMaximumLoot(50, []int{60, 100, 120}, []int{20, 50, 30}), 180.0000)
	assert.Equal(t, getMaximumLoot(10, []int{500}, []int{30}), 166.6667)
}
