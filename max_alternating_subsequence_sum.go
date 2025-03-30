package main 

func maxAlternatingSum(nums []int) int64 {
	if len(nums) < 2 {
		return int64(nums[0])
	}
	prev_sum := 0
	neg_sum := 0
	for _, val := range nums{
		prev_sum = max(prev_sum,  neg_sum + val)
		neg_sum = max(neg_sum, prev_sum - val)
	}
	return int64(prev_sum)
}