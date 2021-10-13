package main

import (
	"fmt"
	"myAlgo/queue"
)

func main()  {
	//a := []int{9,9,8,3}
	//fmt.Println(week01.PlusOne(a))
	//newNode5 := &LinkedList.ListNode{5, nil}
	//newNode4 := &LinkedList.ListNode{4, newNode5}
	//newNodefix := &LinkedList.ListNode{1, newNode4}
	//
	//newNode3 := &LinkedList.ListNode{3, nil}
	//newNode2 := &LinkedList.ListNode{2,newNode3}
	//newNode1 := &LinkedList.ListNode{1,newNode2}

	//newHead := LinkedList.ReverseKGroup(newNode1, 2)
	//newHead := week01.MergeTwoLists(newNode1, newNodefix)
	//for newHead != nil {
	//	fmt.Println(newHead.Val)
	//	newHead = newHead.Next
	//}

	//a := []string{"4","13","5","/","+"}
	//result := stack.EvalRPN(a)
	//fmt.Println(result)
	//
	//fmt.Println(13/5)


	//heights := []int{1,2,3}
	//max := stack.LargestRectangleArea(heights)
	//fmt.Println(max)

	//heights := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	//heights := []int{4,2,0,3,2,5}
	//max := stack.Trap(heights)
	//fmt.Println(max)
	//s := "stst"
	//fmt.Println(string(s[1]))
	//matrix := [][]byte{{49,48,49,48,48},{49,48,49,49,49},{49,49,49,49,49},{49,48,48,49,48}}
	//max := week01.MaximalRectangle(matrix)
	//fmt.Println(max, []byte("1"))
	nums := []int{1,3,-1,-3,5,3,6,7}
	// 3 5 5
	k := 3

	n := 2
	//fmt.Println(queue.MaxSlidingWindow(nums, k))
	fmt.Println(queue.MaxSlidingWindowSlide(nums, k, n))
}
