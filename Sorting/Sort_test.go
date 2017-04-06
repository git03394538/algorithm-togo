package Sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var input = []int{8, 5, 1, 4, 3, 7, 2, 6, 8, 5, 1, 4, 3, 7, 2, 6, 7}
var expect = []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 7, 8, 8}

func TestMergeSort(t *testing.T) {
	assert.Equal(t, expect, MergeSort(input), "they should be equal")
}

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, expect, BubbleSort(input), "they should be eqaul")
}


func TestSelectionSort(t *testing.T) {
	assert.Equal(t, expect, SelectionSort(input), "they should be eqaul")
}

func TestInsertionSort(t *testing.T) {
	assert.Equal(t, expect, InsertionSort(input), "they should be eqaul")
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(input)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(input)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(input)
	}
}
