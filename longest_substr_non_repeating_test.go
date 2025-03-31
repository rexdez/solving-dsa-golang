package main

import "testing"

func TestLongestSubstrNonRepeating(t *testing.T) {
	test_data := []string{
		"tree",
		"babckba",
		"aa",
		"browncolor",
		"b",
		"conneticut",
	}
	want := []int{
		3, 4, 1, 6, 1, 6,
	}

	for i := range test_data {
		got := LongestSubstrNonRepeating(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed.\n Output: %v, Expected: %v test-data: %v", got, want[i], test_data[i])
		}
	}

}