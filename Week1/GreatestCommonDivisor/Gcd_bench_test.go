package GreatestCommonDivisor

import (
	"testing"
	"math/rand"
	"time"
)

var x int
var y int

func init () {
	rand.Seed( time.Now().UnixNano())
	x = rand.Intn(1000000)
	y = rand.Intn(1000000)
}

func BenchmarkSlowGcd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getGcdSlow(x, y)
	}
}

func BenchmarkFastGcd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetGcdFast(x, y)
	}
}