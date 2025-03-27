package main

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test_data := [][]int{
		{10,56,13,57,2,6},
		{36,12,556,90},
	}
	want := [][]int{
		{2,6,10,13,56,57},
		{12,36,90,556},
	}

	for i := range test_data {
		got := quickSort(test_data[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Test Failed - Output: %v | Input: %v | Expected: %v", got, test_data[i], want[i])
		}
	}
}