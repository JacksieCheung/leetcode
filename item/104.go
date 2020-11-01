package main

// 找树最大深度
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(node *TreeNode, count int, max *int) {
	if node.Left != nil {
		dfs(node.Left, count+1, max)
	}
	if node.Right != nil {
		dfs(node.Right, count+1, max)
	}
	if node.Left == nil && node.Right == nil {
		if count > *max {
			*max = count
		}
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max int
	dfs(root, 0, &max)
	return max + 1
}
