package main

import (
	"slices"
)

func largestDivisibleSubset(nums []int) []int {
	slices.Sort(nums)
	var temp []int = make([]int, len(nums))
	var prev_idx []int = make([]int, len(nums))

	var max_idx int
	
	for i := range nums{
		temp[i] = 1
		prev_idx[i] = -1
		for j := range i {
			if nums[i] % nums[j] == 0 && temp[j]+1 > temp[i] {
				temp[i] = temp[j] + 1
				prev_idx[i] = j
			}
		}
		if temp[i] > temp[max_idx]{
			max_idx = i
		}
	}

	// reconstructing the original substring
	var max_subset []int
	for max_idx >=0 {
		max_subset = append([]int{nums[max_idx]}, max_subset...)
		max_idx = prev_idx[max_idx]
	}

	return max_subset
}