package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	n := len(words1)
	m := len(words2)
	if m > n{
		words1, words2 = words2, words1
		m, n = n, m
	}

	var start int
	for start < m && words1[start] == words2[start]{
		start++
	}

	var end int 
	for end < m-start && words1[n-1-end] == words2[m-1-end] {
		end ++
	}
	return start + end == m
}