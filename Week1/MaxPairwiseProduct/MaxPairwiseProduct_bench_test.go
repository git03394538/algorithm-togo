package MaxPairwiseProduct

import (
	"testing"
)
var input = make([]int, 20000)

func BenchmarkSlowMaxPairwiseProduct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMaxPairwiseProductSlow(input)
	}
}

func BenchmarkNormalMaxPairwiseProduct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMaxPairwiseProductNormal(input)
	}
}

func BenchmarkFastMaxPairwiseProduct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getMaxPairwiseProductFast(input)
	}
}