package main

import (
	"fmt"
	"math/rand/v2"
)



func BinaryTreeBuilder(depth int) *TreeNode {
	if depth > 10{
		fmt.Println("Error in creating Tree")
		return nil
	}
	tree := &TreeNode{
		Val: rand.IntN(100),
	}

	growTree(tree, (depth-2))
	return tree
}

func growTree(node *TreeNode, depth int) {
	if depth < 0 {
		return
	}
	// Adding left child
	if randBool() {
		node.Left = &TreeNode{
			Val: rand.IntN(100),
		}
		growTree(node.Left, (depth-1))
	}

	// Growing Right Child
	if randBool() {
		node.Right = &TreeNode{
			Val: rand.IntN(100),
		}
		growTree(node.Right, (depth-1))
	}
}

func randBool() bool {
	a := rand.UintN(10)
	if a < 3 {
		return false
	}
	return true
}

// PrintTree prints the binary tree sideways
func PrintTree(node *TreeNode, prefix string, isLeft bool) {
    if node == nil {
        return
    }

    // Print right child first (so tree grows leftwards)
    if node.Right != nil {
        newPrefix := prefix
        if isLeft {
            newPrefix += "│   "
        } else {
            newPrefix += "    "
        }
        PrintTree(node.Right, newPrefix, false)
    }

    // Print current node
    fmt.Printf("%s", prefix)
    if isLeft {
        fmt.Printf("└── ")
    } else {
        fmt.Printf("┌── ")
    }
    fmt.Printf("%d\n", node.Val)

    // Print left child
    if node.Left != nil {
        newPrefix := prefix
        if isLeft {
            newPrefix += "    "
        } else {
            newPrefix += "│   "
        }
        PrintTree(node.Left, newPrefix, true)
    }
}