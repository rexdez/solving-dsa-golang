package main 

func getLongestAltStr(S string) string{
	n:= len(S)
	if n < 3 {
		return S
	}
	max_str := ""
	start := 0
	for i := range n {
		if i >= 2 && S[i] == S[i-1] && S[i] == S[i-2] {
			start = i-1
		}
		loc:= i-start +1
		if loc > len(max_str) {
			max_str = S[start:i+1]
		}
	}
	return max_str
}