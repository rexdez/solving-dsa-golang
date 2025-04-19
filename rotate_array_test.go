package main

import (
	"math/rand"
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

func TestRotateII(t *testing.T) {
    test_data :=[][]int{
        {2,3,4,5,6},
        {2,3,4,5,6},
        {1,2,3,4,5,6,7},
        {-1,-100,3,99},
    }

    rotate_count := []int{
        5,3,3,3,
    }

    want := [][]int {
        {2,3,4,5,6},
        {4,5,6,2,3},
        {5,6,7,1,2,3,4},
        {-100, 3,99,-1},
    }

    for i := range test_data {
        got := rotateArrII(test_data[i], rotate_count[i])
        if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Test Failed - Output: %v | Input: %v | Expected: %v", got, test_data[i], want[i])
        }
        
    }
}

func BenchmarkRotateArray(b *testing.B) {
    for b.Loop() {
        test_data := make([]int, 0)
        randCount := max(rand.Intn(50), 1)
        for range randCount {
            test_data = append(test_data, max(rand.Intn(100), 1))
        }
        rotateArr(test_data, max(rand.Intn(30), 1))
    }
    
}

func BenchmarkRotateArrayII(b *testing.B) {
    for b.Loop() {
        test_data := make([]int, 0)
        randCount := max(rand.Intn(50), 1)
        for range randCount {
            test_data = append(test_data, max(rand.Intn(100), 1))
        }
        rotateArrII(test_data, max(rand.Intn(30), 1))
    }
    
}