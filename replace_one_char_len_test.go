package main

import "testing"

func TestReplaceOneCharLen(t *testing.T) {
	test_data := []string{
		"baaabbabbb",
		"babba",
		"aa",
		"aaabbaaa",
		// "baaababb",
		// "b",
		// "ab",
		// "aabbaabbaabbaabbaa",
	}
	want := []int{
		7,5,2,6,
		// 0, 8,0, 2,16,
	}

	for i := range test_data {
		got := replaceOneCharLen(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed!\n Output: %v, Expected: %v | Input: %v", got, want[i], test_data[i])
		}
	}
}