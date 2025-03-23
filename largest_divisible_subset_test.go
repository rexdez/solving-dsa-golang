package main

import (
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	test_data := [][]int{
		{1,2,3},
		{1,2,4,8},
		{2,3,7,10,13},
		{2,3,9,27,13},
	}
	want := [][]int{
		{1,2},
		{1,2,4,8},
		{2,10},
		{3,9,27},
	}

	for i, val := range test_data {
		got := largestDivisibleSubset(val)
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Invalid data. Output: %v, Expected: %v \n test-data: %v", got, want, test_data[i])
		}
	}
}