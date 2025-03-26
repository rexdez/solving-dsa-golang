package main

func TrappingRainWater(height []int) int {
	var max_water int
	i :=0
	j := len(height)-1
	left_max := height[i]
	right_max := height[j]

	for i<j {
		if left_max < right_max {
			i++
			left_max = max(left_max, height[i])
			max_water += (left_max - height[i])
		} else {
			j--
			right_max = max(right_max, height[j])
			max_water += (right_max - height[j])
		}
	}
	return max_water
}