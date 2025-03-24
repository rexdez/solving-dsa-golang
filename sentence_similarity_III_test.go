package main

import (
	"testing"
)

func TestAreSentenceSimilar(t *testing.T) {
	test_data1 := []string{
		"My name is Hailey",
		"A lot of words",
		"Eating",
		"Happy birthday to lovely you alisha",
		"FmtdCzv Gp YZf UYJ xD iP tqchblXgqvNVdi",
	}
	test_data2 := []string{
		"My Hailey",
		"of",
		"Eating right now",
		"Happy birthday you alisha",
		"xD iP tqchblXgqvNVdi",
	}

	want := []bool{
		true, false, true, true, true,
	}
	for i := range len(test_data1) {
		got := areSentencesSimilar(test_data1[i], test_data2[i])
		if got != want[i] {
			t.Errorf("Invalid output: %v, Expected: %v | Input: %v", got, want[i], test_data1[i])
		}
	}
}