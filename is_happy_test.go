package main

import (
	"math/rand"
	"testing"
)

func TestIsHappy(t *testing.T) {
	got := IsHappy(9987)
	want := false

	if got != want {
		t.Error("That number should be the opposite")
	}
	
}

func BenchmarkIsHappy(b *testing.B) {
	for b.Loop() {
		IsHappy(rand.Intn(10000))
	}
	
}