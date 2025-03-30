package main

import (
	"reflect"
	"testing"
	"sort"
)
func TestTopKElements(t *testing.T) {
	test_data :=[][]int{
		{1,1,1,2,2,3},
		{1},
		{6,6,6,9,9,12,12,1,1,1,88,88,88,88,88,88},
	}

	k_data := []int {
		2,1,3,
	}
	want := [][]int{
		{1,2},
		{1},
		{1,6,88},
	}

	for i := range test_data {
		got := topKFrequent(test_data[i], k_data[i])
		sort.Ints(got)
		if !reflect.DeepEqual(want[i], got) {
			t.Errorf("Test Failed. Output: %v, Expected: %v \n test-data: %v, %v", got, want[i], test_data[i], k_data[i])
		}
	}

	
}