package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归求解，很简单。
// 求根到叶子节点数字之和
func dfs(root *TreeNode, sum int, nums *[]int) {
	sum = sum*10 + root.Val
	if root.Left != nil {
		dfs(root.Left, sum, nums)
	}
	if root.Right != nil {
		dfs(root.Right, sum, nums)
	}
	if root.Left == nil && root.Right == nil {
		*nums = append(*nums, sum)
		return
	}

}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var nums []int
	dfs(root, 0, &nums)
	var sum int
	fmt.Println(nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}
