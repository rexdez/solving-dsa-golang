package main

import (
	"math/rand"
	"testing"
)

func TestIsHappy(t *testing.T) {
	test_data := []int{
		18382,
		10333,
		932092,
		82811,
		821,
		99,
		032,
		0,
		-1,
		100,
	}

	want := []bool{
		false,
		true,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
	}

	for i := range test_data {
		got := IsHappy(test_data[i])
		if got != want[i] {
			t.Errorf("Test failed: data - %v | output - %v | expected - %v", test_data[i], got, want[i])
		}
	}
}

func BenchmarkIsHappy(b *testing.B) {
	for b.Loop() {
		IsHappy(rand.Intn(10000))
	}
	
}