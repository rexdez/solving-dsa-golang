package main

import "testing"

func TestMaxDepth(t *testing.T) {
	// var data [][]*int = [][]*int{
	// 	{intPtr(1), intPtr(2),intPtr(3),intPtr(4), nil, nil,intPtr(5)},
	// 	{nil},
	// 	{intPtr(3),intPtr(9), intPtr(20),nil,nil,intPtr(15),intPtr(7)},
	// }
	data := []*TreeNode{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		nil,
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
	}

	want := []int{3,2, 0, 3}
	
	for i := range(len(want)) {
		got := (maxDepth(data[i]))
		if got != want[i] {
			t.Errorf("incorrect depth calculation - Output: %d | required: %d", got, want[i])
		}
	}
}