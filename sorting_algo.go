package main


func Merge(arr []int, left, mid, right int){
	
}

func MergeSort(arr []int, left, right int) []int {
	if left < right {
		mid := (left + right)/2

		MergeSort(arr, left, mid)
		MergeSort(arr, mid + 1, right)
		Merge(arr, left, mid, right)
	}
	return arr
}