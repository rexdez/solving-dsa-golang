package main

import (
	"fmt"
)

func BFStraversal() {
	tree := BinaryTreeBuilder(8)

	PrintTree(tree, "", false)
	fmt.Println()
	fmt.Println()

	queue := []*TreeNode{tree}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		
		fmt.Printf("Current Node: %v\n", node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

