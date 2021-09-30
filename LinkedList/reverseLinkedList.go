package LinkedList

func ReverseLinkedList(head *ListNode) *ListNode {
	last := &ListNode{0,nil}
	for head != nil {
		nextHead := head.Next
		last.Next = head
		last = head
		head = nextHead
	}
	return last
}
