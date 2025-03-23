package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestLargestDivisibleSubset(t *testing.T) {
	test_data := [][]int{
		{1,2,3},
		{1,2,4,8},
		{2,3,7,10,13},
		{2,3,9,27,13},
		{4,8,10,240},
	}
	want := [][]int{
		{1,2},
		{1,2,4,8},
		{2,10},
		{3,9,27},
		{4,8,240},
	}

	for i, val := range test_data {
		got := largestDivisibleSubset(val)
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Invalid data. Output: %v, Expected: %v \n test-data: %v", got, want, test_data[i])
		}
	}
}

func BenchmarkLargestDivisibleSubset(b *testing.B) {
	for b.Loop() {
		n :=  rand.Intn(100)
		if n == 0{
			n = 1
		}
		var test_data []int
		for range n {
			var num int
			for num == 0 {
				num = rand.Intn(100)
			}
			test_data = append(test_data, num)
		}
		largestDivisibleSubset(test_data)
	}
}