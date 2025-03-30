package main

import "sort"

type Data struct {
	num int
	freq int
}
func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)
	val_map := make(map[int]int)
	for _, num := range nums {
		val_map[num]++
	}
	counts := make([]Data, 0, len(val_map))
	for key, val := range val_map {
		counts = append(counts, Data{num: key, freq: val})
	}
	sort.Slice(counts,func(i, j int) bool {
		// if counts[i].freq == counts[j].freq {
		// 	return counts[i].num < counts[j].num
		// }
		return counts[i].freq < counts[j].freq
	})
	result := make([]int, 0)
	for _, val := range counts[len(counts)-k:] {
		result = append(result, val.num)
	}
	return result
}