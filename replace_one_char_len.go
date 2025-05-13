package main

func replaceOneCharLen(S string) int{
	max_len := 0
	chars := []rune(S)
	start := 0
	count := 0
	i:=0
	for i < len(chars){
		if i>=2 && chars[i] == chars[i-1] && chars[i] == chars[i-2]{
			count ++
			if count < 2 {
				if chars[i] == 'a' {
					chars[i] = 'b'
				} else {
					chars[i] = 'a'
				}
			} else {
				max_len = max(max_len, i-start+1)
				chars = []rune(S)
				start++
				i=start
				count =0
			}
		}
		max_len = max(max_len, i-start+1)
		i++
	}
	return max_len
}