package main

import (
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	test_data := []string{
		"tree", "cccaaa", "Aabb",
	}

	want := []string{
		"eert", "aaaccc", "bbAa",
	}

	for i := range test_data {
		got := frequencySort2(test_data[i])
		if got != want[i] {
			t.Errorf("Invalid data. Output: %v, Expected: %v \n test-data: %v", got, want[i], test_data[i])
		}
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func BenchmarkFreqSort(b *testing.B) {
	for b.Loop() {
		frequencySort(RandStringBytes(rand.Intn(20)))
	}
}
func BenchmarkFreqSort2(b *testing.B) {
	for b.Loop() {
		frequencySort2(RandStringBytes(rand.Intn(20)))
	}
}