package main 

func quickSort(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}

	pivot :=  arr[0]
	var left, right []int 
	for _, val := range arr[1:] {
		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	res := append(append(left, pivot), right...)
	return res
}