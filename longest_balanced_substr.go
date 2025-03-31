package main

// import "fmt"

func longestBalancedStr(S string) int {
	balance := 0
	maxLen := 0
	indexMap := make(map[int]int)
	indexMap[0] = -1 // Handle balance=0 from start

	for i := range S {
		if S[i] == 'a' {
			balance++
		} else {
			balance--
		}

		if startIdx, found := indexMap[balance]; found {
			maxLen = max(maxLen, i - startIdx)
		} else {
			indexMap[balance] = i
		}
	}

	return maxLen
}