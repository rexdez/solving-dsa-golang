package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// Finding minimum of average of minimum and maximum value in an array
func main() {
	fmt.Println(addTwoNumbers(
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: nil,
					},
				},
			},
		},
		&ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: nil,
			},
		},
	))
}

