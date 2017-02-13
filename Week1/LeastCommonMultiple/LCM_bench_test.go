package LeastCommonMultiple

import (
	"testing"
	"math/rand"
	"time"
)

var x int
var y int

func init () {
	rand.Seed( time.Now().UnixNano())
	x = rand.Intn(50)
	y = rand.Intn(50)
}

func BenchmarkSlowLCM(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getLCMSlow( x, y)
	}
}

func BenchmarkFastLCM(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getLCMFast( x, y)
	}
}
