package LinkedList

func ReverseKGroup(head *ListNode, k int) *ListNode {
	protect := &ListNode{0, head}
	last := protect

	for head != nil {
		// 分组
		end := getEnd(head, k)
		if end == nil {
			break
		}

		// 找出下一组节点的头结点
		nextGroupHead := end.Next

		// 翻转长度为k的链表
		myReverse(head, nextGroupHead)

		// 处理边
		last.Next = end
		head.Next = nextGroupHead

		last = head
		head = nextGroupHead
	}

	return protect.Next
}

func getEnd(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}

		head = head.Next
	}
	return nil
}

func myReverse(head *ListNode, end *ListNode) {
	last := head
	head = head.Next
	for head != end {
		nextHead := head.Next

		head.Next = last
		last = head
		head = nextHead
	}
}