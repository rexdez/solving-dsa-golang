package main

import "testing"

func TestGetLongestAltStr(t *testing.T) {
	test_data := []string{
		"baaabbabbb",
		"babba",
		"aa",
		"aaa",
		"b",
		"aabbaaaabbaabbaa",
	}
	want := []string{
		"aabbabb",
		"babba",
		"aa",
		"aa",
		"b",
		"aabbaabbaa",
	}

	for i := range test_data {
		got := getLongestAltStr(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed!\n Output: %v, Expected: %v | Input: %v", got, want[i], test_data[i])
		}
	}
}