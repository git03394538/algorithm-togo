package DP2_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"Algorithm/Week5/DP2"
)

//http://www.programcreek.com/2013/12/edit-distance-in-java/
func TestEditDistance(t *testing.T) {
	assert.Equal(t,
		8,
		DP2.EditDistance("dynamic", "programming"))
}

//DP2 is very similar with Longest Common Subsequence
//http://www.cs.usfca.edu/~galles/visualization/DPLCS.html
func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t,
		3,
		DP2.LongestCommonSubsequence("dynamic", "programming", 6, 10))
}