// 2457. Minimum Addition to Make Integer Beautiful
package main

func MakeIntegerBeautiful(n int64, target int) int64 {
	if getDigitSum(n) <= target {
		return 0
	}
	base := int64(1)
	num := n
	for getDigitSum(n) > target {
		n = (n/10) + 1
		base *=10

	}
	return (n*base) - num
}

func getDigitSum(num int64) int {
	var sum int = 0
	for num > 0 {
		sum += int(num%10)
		num /= 10
	}
	return sum
}