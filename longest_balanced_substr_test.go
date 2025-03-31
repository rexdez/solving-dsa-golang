package main

import "testing"

func TestLongestBalancedStr(t *testing.T) {
	test_data := []string{
		"baaabbabbb",
		"babba",
		"aa",
		"baaababb",
		"b",
		"ab",
		"aabbaabbaabbaabbaa",
	}
	want := []int{
		8, 4, 0, 8,0, 2,16,
	}

	for i := range test_data {
		got := longestBalancedStr(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed!\n Output: %v, Expected: %v | Input: %v", got, want[i], test_data[i])
		}
	}
}