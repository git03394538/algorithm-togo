package LeastCommonMultiple

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlowGcb(t *testing.T) {
	assert.Equal(t, getLCMSlow(6, 8), 24)
}

func TestFastGcb(t *testing.T) {
	assert.Equal(t, getLCMFast(6, 8), 24)
	assert.Equal(t, getLCMFast(28851538, 1183019), 1933053046)
}

