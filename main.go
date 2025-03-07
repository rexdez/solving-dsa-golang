package main

import (
	"fmt"
)

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
	fmt.Println(isSameTree(
		&TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 9,
				Left: nil,
			},
			Right: &TreeNode{
				Val: 8,
				Right: nil,
			},
		},
		&TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 8,
				Right: nil,
			},
		},
	))
}

