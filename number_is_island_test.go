package main

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	defer func() {
		if r := recover(); r!=nil{
			fmt.Println("Panic Error: ", r)
		}
	}()
	test_data := [][][]byte{
		{
			{'1','1','1','1','0'},
			{'1','1','0','1','0'},
			{'1','1','0','0','0'},
			{'0','0','0','0','0'},
		},
	}
	want := []int{
		1,
	}

	for i := range test_data {
		got := numIslands(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed: Output: %v | Expected: %v | Data: %v", got, want[i], test_data[i])
		}
	}
}