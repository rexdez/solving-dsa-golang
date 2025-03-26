package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.

// Open brackets must be closed in the correct order.

// Every close bracket has a corresponding open bracket of the same type.

// "[]({})" - valid

// "[]({)}" - invalid
// "[]({}" - invalid

// "{[()]}" valid

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