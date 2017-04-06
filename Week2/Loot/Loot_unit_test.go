package Loot_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"Algorithm/Week2/Loot"
)

func TestGetMaximumLoot(t *testing.T) {
	assert.Equal(t, 180.0000, Loot.GetMaximumLoot(50, []int{60, 100, 120}, []int{20, 50, 30}))
	assert.Equal(t, 166.6667, Loot.GetMaximumLoot(10, []int{500}, []int{30}))
}
