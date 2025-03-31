package main


func LongestSubstrNonRepeating(S string) int {
	n := len(S)
	if n < 1 {
		return n
	}
	max_len := 0
	start := 0
	char_map := make(map[byte]int)
	for i := range S {
		if idx, ok := char_map[S[i]]; ok && idx >= start{
			start = idx + 1
		}
		char_map[S[i]] = i
		max_len = max(max_len, i-start + 1)
	}
	return max_len
}