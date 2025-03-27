package main

import (
	"reflect"
	"testing"
)


func TestRotate(t *testing.T) {
    test_data :=[][]int{
        {2,3,4,5,6},
        {1,2,3,4,5,6,7},
        {-1,-100,3,99},
    }

    rotate_count := []int{
        3, 3,3,
    }

    want := [][]int {
        {4,5,6,2,3},
        {5,6,7,1,2,3,4},
        {-100, 3,99,-1},
    }

    for i := range test_data {
        got := test_data[i]
        rotateArr(got, rotate_count[i])
        if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Test Failed - Output: %v | Input: %v | Expected: %v", got, test_data[i], want[i])
        }
        
    }
}