package main

// 给定一个二叉树，检查它是否是镜像对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var isSameTree func(R *TreeNode, L *TreeNode) bool
	isSameTree = func(R *TreeNode, L *TreeNode) bool {
		if R == nil && L == nil {
			return true
		} else if R == nil && L != nil {
			return false
		} else if R != nil && L == nil {
			return false
		}
		if R.Val == L.Val {
			return isSameTree(R.Left, L.Right) && isSameTree(R.Right, L.Left)
		}
		return false
	}

	return isSameTree(root.Left, root.Right)
}
