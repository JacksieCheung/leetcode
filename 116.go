package main

// 填充每个节点下一个右侧节点指针
// 通过二维数组把每个节点弄出来存储，时间复杂度是O(n+i*j)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	var nums [][]*Node
	dfs(root, 0, &nums)
	//fmt.Println(*(nums[2][3]))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i])-1; j++ {
			nums[i][j].Next = nums[i][j+1]
		}
		nums[i][len(nums[i])-1] = nil
	}
	return root
}

func dfs(node *Node, depth int, nums *[][]*Node) {
	if node == nil {
		return
	}

	if depth+1 > len(*nums) {
		*nums = append(*nums, []*Node{})
	}
	(*nums)[depth] = append((*nums)[depth], node)

	dfs(node.Left, depth+1, nums)
	dfs(node.Right, depth+1, nums)
}
