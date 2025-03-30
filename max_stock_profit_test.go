package main

import "testing"

func TestMaxStockProfit(t *testing.T) {
	test_data := [][]int{
		{7,1,5,3,6,4},
		{1,2,3,4,5},
		{7,6,4,3,1},
	}
	want := []int{
		7,4,0,
	}
	for i := range test_data {
		got := maxProfit(test_data[i])
	if got != want[i] {
		t.Errorf("Test failed for \ndata at: %v | Expected: %v | Output: %v",i+1, want[i], got)
	}
	}
}