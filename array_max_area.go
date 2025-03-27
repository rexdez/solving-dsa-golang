package main

func maxArrayArea(height []int) int {
	var max_area int
	i := 0
	j := len(height) -1 
	for i < j{
		max_area = max(max_area, (j-i)*min(height[i], height[j]))
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max_area
}
