package main

import (
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	test_data := [][]int{
		{5, 2, 8, 6, 3, 6, 9, 5},
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
		{1},
	}
	want := []int{4, 4, 4, 1, 1}

	for i, val := range test_data {
		got := lengthOfLIS(val)
		if got != want[i] {
			t.Errorf("Invalid output: %v, required: %v, input: %v", got, want[i], val)
		}
	}
}
