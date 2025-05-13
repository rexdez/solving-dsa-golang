package main


func IsHappy(n int) bool {
	dump := make(map[int]bool)
	for n != 1 {
		n = getSumOfSquares(n)
		if _, ok := dump[n]; ok {
			return false
		}
		dump[n] = false
	}
	return n == 1
}

func getSumOfSquares(n int) int {
	var sum int
	for n > 0 {
		a := n % 10
		sum += (a * a)
		n /= 10
	}
	return sum
}