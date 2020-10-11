package main

// 给出两个 非空 的链表用来表示两个非负的整数。
// 其中，它们各自的位数是按照 逆序 的方式存储的，
// 并且它们的每个节点只能存储 一位 数字。
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	result := l3
	flag := 0
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + flag
		if tmp > 9 {
			node := &ListNode{tmp % 10, nil}
			l3.Next = node
			l3 = l3.Next
			l1 = l1.Next
			l2 = l2.Next
			flag = 1
		} else {
			node := &ListNode{tmp, nil}
			l3.Next = node
			l3 = l3.Next
			l1 = l1.Next
			l2 = l2.Next
			flag = 0
		}
	}
	for l1 != nil {
		tmp := l1.Val + flag
		if tmp > 9 {
			node := &ListNode{tmp % 10, nil}
			l3.Next = node
			l3 = l3.Next
			l1 = l1.Next
			flag = 1
		} else {
			node := &ListNode{tmp, nil}
			l3.Next = node
			l3 = l3.Next
			l1 = l1.Next
			flag = 0
		}
	}
	for l2 != nil {
		tmp := l2.Val + flag
		if tmp > 9 {
			node := &ListNode{tmp % 10, nil}
			l3.Next = node
			l3 = l3.Next
			l2 = l2.Next
			flag = 1
		} else {
			node := &ListNode{tmp, nil}
			l3.Next = node
			l3 = l3.Next
			l2 = l2.Next
			flag = 0
		}
	}
	if flag == 1 {
		node := &ListNode{
			Val:  1,
			Next: nil,
		}
		l3.Next = node
	}
	return result.Next
}
