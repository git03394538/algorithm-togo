package DPChange_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"Algorithm/Week5/DPChange"
)

func TestChange(t *testing.T) {
	assert.Equal(t, 0, DPChange.RecursiveChange(0, []int{1, 5, 10}))
	assert.Equal(t, 5, DPChange.RecursiveChange(9, []int{1, 5, 10}))

	assert.Equal(t, 0, DPChange.DPChange(0, []int{1, 5, 10}))
	assert.Equal(t, 5, DPChange.DPChange(9, []int{1, 5, 10}))
}