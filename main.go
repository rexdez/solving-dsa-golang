package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// Finding minimum of average of minimum and maximum value in an array
func main() {
	obj := Constructor(3, []int{1,4,5,8,2})
	fmt.Println(obj)
 	param_1 := obj.Add(3)
	fmt.Println(param_1)
}