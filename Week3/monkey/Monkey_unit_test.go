package monkey_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"Algorithm/Week3/monkey"
)

func TestMonkey(t *testing.T) {
	assert.Equal(t, 4 , monkey.Do(2, 2, 1, 1))
	assert.Equal(t, 1534 , monkey.Do(10, 2, 1, 1))
	assert.Equal(t, 12203745396 , monkey.Do(20, 3, 5, 3))
}