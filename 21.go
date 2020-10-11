package main

// 将两个升序链表合并为一个新的 升序 链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。
// 两条链表各走一遍，时间复杂度o(n+m)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res ListNode
	p := &res

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	for l1 != nil {
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}
	return res.Next
}
