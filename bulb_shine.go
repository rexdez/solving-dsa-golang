package main

func calcShines(a []int) int {
	bufferOn := make([]int, len(a))
	shineCount := 0
	for _, val:= range a {
		bufferOn[val-1] = 1

		j := 0
		for j < val {
			if bufferOn[j] != 1 {
				break
			}
			j++
		}
		if j == val {
			shineCount++
		}
	}
	return shineCount
}