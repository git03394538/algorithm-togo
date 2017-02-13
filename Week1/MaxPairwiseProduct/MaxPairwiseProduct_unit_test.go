package MaxPairwiseProduct

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
)

var inputLargeNumber = []int{1, 9000000000, 100000000, 2, 134345, 92345}
var inputRandNumber = []int{}

func init() {
	rand.Seed( time.Now().UnixNano())
	for i := rand.Intn(10000) ; i >= 0 ; i-- {
		inputRandNumber = append(inputRandNumber, rand.Intn(10000000))
	}
}

func TestSlowMaxPairwiseProduct(t *testing.T) {
	assert.Equal(t, int64(900000000000000000),
		getMaxPairwiseProductSlow(inputLargeNumber),
		"they should be equal")
}

func TestNormalMaxPairwiseProduct(t *testing.T) {
	assert.Equal(t, int64(900000000000000000),
		getMaxPairwiseProductNormal(inputLargeNumber),
		"they should be equal")
}

func TestFastMaxPairwiseProduct(t *testing.T) {
	assert.Equal(t, int64(900000000000000000),
		getMaxPairwiseProductFast(inputLargeNumber),
		"they should be equal")
}

func TestThreeMaxPairwiseProductReturnTheSame(t *testing.T) {
	slowResult := getMaxPairwiseProductSlow(inputRandNumber)
	normalResult := getMaxPairwiseProductNormal(inputRandNumber)
	fastResult := getMaxPairwiseProductFast(inputRandNumber)
	assert.Equal(t, slowResult, normalResult, "slow and normal should be equal")
	assert.Equal(t, normalResult, fastResult, "normal and fast should be equal")
}
