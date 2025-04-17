package main

import "testing"

func TestBinaryGap(t *testing.T) {
	test_data := []int {
		590043,
		50092,
		129,
		17,
		90001,
		2305,
		1152,
	}

	want := []int{8,4,6,3,3,7,2}

	for i := range test_data {
		got := binaryGap(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed Output: %v, Expected: %v | Input: %v", got, want[i], test_data[i])
		}
	}

}