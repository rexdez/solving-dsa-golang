package main


func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	var depth1, depth2 int
	depth1 = maxDepth(root.Left)
	depth2 = maxDepth(root.Right)
	return max(depth1, depth2) + 1
}

func max(a, b int) int {
	if a > b{
		return a 
	}
	return b
}