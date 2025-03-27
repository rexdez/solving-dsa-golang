package main


import (
	"testing"
)

func TestIsValidString(t *testing.T) {
	test_data := []string{
		"[]({})", "[]({)}", "[]({}", "{[()]}",
	}
	want := []bool{
		true, false, false, true,
	}

	for i := range test_data {
		got := isValidString(test_data[i])
		if got != want[i] {
			t.Errorf("Test Failed- Output: %v | Expected: %v | Data: %v", got, want[i], test_data[i])
		}
	}
}