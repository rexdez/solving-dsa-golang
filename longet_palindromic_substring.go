package main

import (
)

func LongestPalindrome(s string) string {
	longest_palindrome := ""
	for i := range len(s) {
		pal_str := getPalindrome(s, i, i)
		pal_str2 := getPalindrome(s, i, i+1)

		if len(pal_str) > len(longest_palindrome) {
			longest_palindrome = pal_str
		}
		if len(pal_str2) > len(longest_palindrome) {
			longest_palindrome = pal_str2
		}
	}
	return longest_palindrome
}

func getPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1:r]
}