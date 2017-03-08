package MaximumRevenue

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMaximumRevenue(t *testing.T) {
	assert.Equal(t, 897, getMaximumRevenue(1, []int{23}, []int{39}))
	assert.Equal(t, 23, getMaximumRevenue(3, []int{1, 3, -5}, []int{-2, 4, 1}))
}