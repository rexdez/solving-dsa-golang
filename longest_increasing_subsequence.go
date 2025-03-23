package main

func lengthOfLIS(nums []int) int {
	max_len := make([]int, len(nums))
	for i := range nums {
		max_len[i] = 1
	}
	var maxres int = 1
	for i := 1; i < len(nums); i++ {
		for j := range i {
			if nums[j] < nums[i] {
				max_len[i] = max(max_len[i], max_len[j]+1)
			}
		}
		maxres = max(max_len[i], maxres)
	}
	return maxres
}
