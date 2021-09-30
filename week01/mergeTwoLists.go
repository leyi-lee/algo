package week01

import "myAlgo/LinkedList"

func MergeTwoLists(l1 *LinkedList.ListNode, l2 *LinkedList.ListNode) *LinkedList.ListNode {
	head  := &LinkedList.ListNode{0, nil}
	newNode := head
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val <= l2.Val) {
			newNode.Next = l1
			l1 = l1.Next
		} else {
			newNode.Next = l2
			l2 = l2.Next
		}
		newNode = newNode.Next
	}
	return head.Next
}
