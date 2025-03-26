package main 

import (
    "testing"
)

func TestIntToRoman(t *testing.T) {
    data := []int{
        3749,
        1991,
    }
    want := []string{
        "MMMDCCXLIX",
        "MCMXCI",
    }

    for i := range data {
        got := intToRoman(data[i])
        if got != want[i] {
			t.Errorf("Test Failed Output: %v, Expected: %v | Input: %v", got, want[i], data[i])
        }
    }
     
}