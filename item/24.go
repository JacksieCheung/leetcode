package main

// 两两交换链表节点。
// 给的是单链表，交换节点有三个指针要变：
// pre->cur->next->next2
// curNode.Next=nextNode.Next,/ curNode.Next=curNode.Next.Next
// nextNode.Next=curNode,
// preNode.Next=nextNode.
// 因为没有前驱指针，所以要用两个指针分别指向cur和next来操作

// 双向
// node.Next.Pre=w
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	ptr := head
	sol(&head, ptr, nil)
	return head
}

func sol(head **ListNode, node *ListNode, preNode *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	nexNode := node.Next
	node.Next = nexNode.Next
	nexNode.Next = node
	if preNode != nil {
		preNode.Next = nexNode
	} else {
		*head = nexNode
	}
	sol(nil, node.Next, node)
}

// 牺牲代码可读性做的空间优化，少声明了一个变量。
// 先用pre 保存 next 值，再变 node.Next　再用 pre 访问　原来的 nodeNext
func sol(head **ListNode, node *ListNode, preNode *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	if preNode != nil {
		preNode.Next = node.Next
	} else {
		*head = node.Next
	}
	node.Next = node.Next.Next
	if preNode != nil {
		preNode.Next.Next = node
	} else {
		(*head).Next = node
	}
	sol(nil, node.Next, node)
}
