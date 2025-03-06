package main

import "fmt"

func IsHappy(n int) bool {
	for n != 1 && n != 4 {
		n = getSumOfSquares(n)
		fmt.Println(n)
	}
	return n == 1
}

func getSumOfSquares(n int) int {
	var sum int
	for n > 0 {
		a := int((n % 10))
		sum +=  (a*a)
		n /= 10
	}
	return sum
}