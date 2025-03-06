package main

func AddDigits(n int) int {
	for n >= 10 {
		n = getSum(n)
	}
	return n
}

func getSum(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}