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
	ch := make(chan int)
	ch2 := make(chan chan int)
	go func () {
		ch2 <- ch
	}()
	
	res := <- ch2
	go func() {
		ch <- 555
	}()
	val := <- res
	fmt.Println(val)
}

