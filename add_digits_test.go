package main

import (
	"math/rand"
	"testing"
)

func TestAddDigits(t *testing.T) {
	got := AddDigits(5663)
	want := 2
	if got != want {
		t.Errorf("Invalid opertion - Output: %d | Expected: %d", got , want)
	}
}

func BenchmarkAddDigits(b *testing.B) {
	for b.Loop() {
		AddDigits(rand.Intn(10000))
	}
}