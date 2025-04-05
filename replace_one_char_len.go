package main

import "fmt"

func replaceOneCharLen(S string) int{
	max_len := 0
	chars := []rune(S)
	start := 0
	count := 0
	i:=0
	for i < len(chars){
		if i>=2 && chars[i] == chars[i-1] && chars[i] == chars[i-2]{
			fmt.Println("  ",string(chars[i]), string(chars[i-1]), string(chars[i-2]))
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
		fmt.Println(start, string(chars))
		max_len = max(max_len, i-start+1)
		i++
	}
	return max_len
}