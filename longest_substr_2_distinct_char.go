package main

func LongestSubstr2DistinctChar(S string) int {
	n := len(S)
	if n < 3 {
		return n
	}
	max_len := 0
	char_map := make(map[byte]int)
	start := 0
	for i := range n {
		char_map[S[i]]++
		for len(char_map) > 2{
			char_map[S[start]]--
			if char_map[S[start]] == 0 {
				delete(char_map, S[start])
			}
			start++
		}
		max_len = max(max_len, i-start+1)
	}
	return max_len
}

func LongestSubstr2DistinctChar2(S string) int {
	n := len(S)
	if n < 3 {
		return n
	}
	max_len := 0
	char_map := make(map[byte]int)
	start := 0
	i := 0
	for i < len(S) {
		if _, ok := char_map[S[i]]; !ok  && len(char_map) == 2{
			temp_len := 0
			for _, val := range char_map{
				temp_len += val
			}
			max_len = max(max_len, temp_len)
			i = start
			char_map = map[byte]int{}
			continue
		} else if !ok  && len(char_map) < 2{
			start = i
		}
		char_map[S[i]]++
		i++
	}
	temp_len := 0
	for _, val := range char_map{
		temp_len += val
	}
	max_len = max(max_len, temp_len)
	return max_len
}