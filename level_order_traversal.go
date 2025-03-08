package main

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var depth int = 0
	var arr [][]int
	arr = append(arr, []int{root.Val})
	arr = addNodes(arr, depth+1, root.Left, root.Right)
	return arr
}

func addNodes(arr [][]int, depth int, left, right *TreeNode) [][]int {
	if (left != nil || right != nil) && len(arr) <= depth {
		arr = append(arr, []int{})
	}
	if left != nil {
		arr[depth] = append(arr[depth], left.Val)
		arr = addNodes(arr, depth+1, left.Left, left.Right)
	}
	if right != nil {
		arr[depth] = append(arr[depth], right.Val)
		arr = addNodes(arr, depth+1, right.Left, right.Right)
	}
	return arr
}