package main

// 给你一棵所有节点为非负值的二叉搜索树，
// 请你计算树中任意两节点的差的绝对值的最小值。
// 第一种方法，遍历存数组再在数组查找，时间复杂度o(m+n2)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nums []int

func getMinimumDifference(root *TreeNode) int {
	nums = nil
	dfs(root)
	fmt.Println(len(nums))
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}
	fmt.Println(nums)
	minlen := count(nums[0], nums[1])
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if count(nums[i], nums[j]) < minlen {
				minlen = count(nums[i], nums[j])
			}
		}
	}
	return minlen
}

func dfs(root *TreeNode) {
	nums = append(nums, root.Val)
	if root.Left != nil {
		dfs(root.Left)
	}
	if root.Right != nil {
		dfs(root.Right)
	}
	return
}

func count(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
