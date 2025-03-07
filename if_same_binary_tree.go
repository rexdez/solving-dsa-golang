package main

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	fmt.Println("p:", p, "\nq:", q)
	if p == nil && q ==nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
} 