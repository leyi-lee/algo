package LinkedList

func reverseLinkedList(head *ListNode) *ListNode {
	last := &ListNode{0,nil}
	for head != nil {
		nextHead := head.Next
		last.Next = head
		last = head
		head = nextHead
	}
	return last
}
