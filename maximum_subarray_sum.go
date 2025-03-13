package main

func MaxSubArray(nums []int) int {
	var largest_sum int = nums[0]
    temp_sum:= 0
	for _, val := range nums {
		if temp_sum < 0 {
            temp_sum = 0
        }
        temp_sum += val
        if temp_sum > largest_sum {
            largest_sum = temp_sum
        }
    }
	return largest_sum
}
