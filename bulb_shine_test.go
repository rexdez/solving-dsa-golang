package main

import (
	"reflect"
	"testing"
)

func TestCalcShines(t *testing.T) {
	test_data := [][]int{
		{2, 3, 5, 1, 4},
		{2, 3, 1, 4, 5},
	}

	want := []int{
		2, 3,
	}

	for i := range test_data {
		got := calcShines(test_data[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Test Failed: data: %v | got: %v | expected: %v", test_data[i], got, want[i])
		}
	}
}
