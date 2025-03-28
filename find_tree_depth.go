package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth1, depth2 int
	depth1 = maxDepth(root.Left)
	depth2 = maxDepth(root.Right)
	return max(depth1, depth2) + 1
}

// This is primarily DFS not BFS
func MaxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depth1, depth2 int
	depth1 = MaxDepthBFS(root.Left)
	depth2 = MaxDepthBFS(root.Right)
	return getmax(depth1, depth2) + 1
}

func getmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
