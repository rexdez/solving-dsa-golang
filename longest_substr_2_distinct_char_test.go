package main

import (
	"math/rand"
	"testing"
)

func TestLongestSubstr2DistinctChar(t *testing.T) {
	test_data := []string{
		"tree",
		"bbabckba",
		"aa",
		"abbaaabba",
		"bkkkkkkkkk",
		"co",
		"rszdxtttttttttttttttttaaaaaaaaaaaaaaaaaaaaaafcgvh",
	}
	want := []int{
		3, 4, 2, 9, 10, 2, 39,
	}

	for i := range test_data {
		got := LongestSubstr2DistinctChar(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed.\n Output: %v, Expected: %v test-data: %v", got, want[i], test_data[i])
		}
	}

}

var short_charset = "abcdef"

func RandSmallStringBytes(n int, charset string) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}
func BenchmarkLongestSubstr2DistinctChar(b *testing.B) {
	for b.Loop() {
		LongestSubstr2DistinctChar(RandSmallStringBytes(rand.Intn(100), short_charset))
	}
}
