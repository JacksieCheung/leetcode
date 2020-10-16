package main

// 广度优先遍历
// 通过队列实现，迭代运行，每次迭代把当前节点的左右节点放进队列，下次迭代取出处理
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	// 初始化队列同时将第一层节点加入队列中，即根节点
	queue := []*Node{root}

	// 循环迭代的是层数
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		// 遍历这一层的所有节点
		for i, node := range tmp {
			// 连接
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 拓展下一层节点
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	// 返回根节点
	return root
}
