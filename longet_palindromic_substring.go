package main

import (
)
// jjabakabaccd
func longestPalindrome(s string) string {
	var longest_str string
	chars := []rune(s)
	for i:=0; i < len(s);i++{
		for j:=i+1; j < len(s); j++ {
			str := string(chars[i:j+1])
			if chars[j] == chars[i] && str == reverseStr(str){
				longest_str == str
			}
		}
	}
	return longest_str
}

func reverseStr(s string) string{
	chars := []rune(s)
	for i, j := 0,  len(s); j > i; i,j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}