package main


func rotateArrII(nums []int, k int) []int{
    n := len(nums)
	loc := n-(k%n)
	
	res := nums[loc:]
	res = append(res, nums[:loc]...)
	return res
}


func rotateArr(nums []int, k int) {
    n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, k, n-1)
	reverse(nums, 0, k-1)
}

func reverse(nums []int, m, n int) {
	for i, j := m, n; i<j; i,j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}