package ChangingMoney

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestChangingMoney(t *testing.T) {
	assert.Equal(t, change(2), 2)
	assert.Equal(t, change(28), 6)
}
