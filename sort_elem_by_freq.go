package main

import (
	"container/heap"
	"sort"
)

type Char struct {
    char rune
    freq int
}

func frequencySort(s string) string {
	char_map := make(map[rune]int)
	for _, val := range s {
		char_map[val]++
	}
	list := make([]Char, 0, len(char_map))
	for key, val := range char_map {
		list = append(list, Char{char: key, freq: val})
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].freq == list[j].freq {
			return list[i].char < list[j].char
		}
		return list[i].freq > list[j].freq
	})

	result := ""
	for _, val := range list{
		for range val.freq {
			result += string(val.char)
		}
	}
	return result
    
}

type CharList []*Char

func (c CharList) Len() int {
	return len(c)
}

func (c CharList) Less(i, j int) bool{
	if c[i].freq == c[j].freq {
		return c[i].char < c[j].char

	}
	return c[i].freq > c[j].freq
}

func (c *CharList) Pop() any{
	old := *c
	res := old[(old.Len()-1)]
	*c = old[:(old.Len()-1)]
	return res
}

func (c *CharList) Push(i any){
	*c = append(*c, i.(*Char))
}

func (c CharList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}


func frequencySort2(s string) string {
	cl := &CharList{}
	heap.Init(cl)

	char_map := make(map[rune]int)
	for _, val := range s {
		char_map[val]++
	}
	for key, val := range char_map {
		heap.Push(cl, &Char{
			char: key,
			freq: val,
		})
	}
	result := make([]rune, 0)
	for cl.Len() > 0{
		val := heap.Pop(cl).(*Char)
		for range val.freq {
			result = append(result, val.char)
		}
	}
	return string(result)
}