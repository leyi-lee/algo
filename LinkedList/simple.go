package LinkedList

type ListNode struct {
	Val int
	Next *ListNode
}

func Merge(nums1 []int, m int, nums2 []int, n int)  {
	//i := m-1
	//j := n-1
	//
	//for k:= m+n-1;k >= 0;k-- {
	//	if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
	//		nums1[k] = nums1[i]
	//		i--
	//	} else if j >= 0 {
	//		nums1[k] = nums2[j]
	//		j--
	//	}
	//}

	tmp := make([]int, m+n)
	i := 0
	j := 0
	for k := 0; k <= m+n-1; k++ {
		if i < m && j < n && nums1[i] < nums2[j] {
			tmp[k] = nums1[i]
			i++
		} else if j < n {
			tmp[k] = nums2[j]
			j++
		}
	}

	for t := 0; t < len(tmp); t++ {
		nums1[t] = tmp[t]
	}
}

func ReverseSingleLinked(headNode *ListNode) *ListNode {
	var last *ListNode
	for headNode != nil {
		nextHead := headNode.Next
		headNode.Next = last
		last = headNode
		headNode = nextHead
	}
	return last
}



