package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name *string
}

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
	//nums := []int{1,3,-1,-3,5,3,6,7}
	// 3 5 5
	//k := 3

	//n := 2
	//fmt.Println(queue.MaxSlidingWindow(nums, k))
	//fmt.Println(queue.MaxSlidingWindowSlide(nums, k, n))
	//fmt.Println(hash.TwoSum([]int{3,2,4}, 6))
	//fmt.Println(PrefixSum.NumberOfSubArrays([]int{2,2,2,1,2,2,1,2,2,2}, 2))

	//b := 10
	//a := make([][10]int, b)
	//fmt.Println(a[0][0])
	//fmt.Println(len(a), cap(a))
	/**
		1 1 1 1    1,2,3,4
		1 1 1 1	   2,4,6,8
		1 1 1 1	   3,6,9,12
	 */
	//sums := PrefixSum.TwoDimPreSum([][]int{{1,1,1,1},{1,1,1,1},{1,1,1,1}})
	//fmt.Println(sums[3][4])

	//sums := PrefixSum.TwoDimPreDiff([][]int{{1,1,1,1},{1,1,1,1},{1,1,1,1}})
	//fmt.Println(sums[3][4])

	//fmt.Println(PrefixSum.PreDiff([]int{1,2,3,4,5,6}))
	//matrix := [][]int{{0,1,0},{1,1,1},{0,1,0}}
	//result := week02.NumSubmatrixSumTarget(matrix, 0)
	//fmt.Println(result)

	//a := []int{1,2,0,4,5,1,2,3}
	//
	//pre := 0
	//k := 12
	//sums := map[int]int{0:1}
	//ans := 0
	//for i:=1;i <= len(a);i++ {
	//	pre += a[i - 1]
	//	if _, ok := sums[pre-k];ok {
	//		ans += sums[pre-k]
	//	}
	//
	//	sums[pre] += 1
	//}
	//fmt.Println(sums)
	//fmt.Println(ans)

	//lru := week02.Constructor(2)
	// [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	//lru.Get(2)
	//lru.Put(2,6)
	//
	//lru.Get(1)
	//
	//lru.Put(1,5)
	//
	//lru.Put(1,2)
	//
	//lru.Get(1)
	//lru.Get(2)
	//lru.Print()



	//lru.Print()


	//fmt.Println(lru.Get(2))

	//lru.Put(4,4)
	//
	//lru.Get(1)
	//
	//lru.Get(3)
	//lru.Get(4)

	//lru.Print()
	//inorder := []int{9,3,15,20,7}
	//postorder := []int{9,15,7,20,3}
	//week03.BuildTree(inorder, postorder)
	//prerequisites := [][]int{{1,0},{1,2},{0,1}}
	//result := week03.FindOrder(3, prerequisites)
	//fmt.Println(result)
	//s := "23"
	//a := s[:len(s) - 1]
	//fmt.Println(s, string(s[1]), len(s), a)
	//piles := []int{1000000000}
	//H := 2
	//a := week05.MinEatingSpeed(piles, H)
	//fmt.Println(a)
	//nums := []int{2,4,5,1,3}
	//dsort.MergeSort(nums, 0, 4)
	//fmt.Println(nums)
	//nums := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	//fmt.Println(week05.SearchMatrix(nums, 3))
	//nums := []int{1,0,-1,0,-2,2}
	//target := 0
	//recur.FourSum(nums, target)
	//nums := []int{1,1,1,1}
	//nums := make([]int, 2, 5)
	//nums[1] =  1
	//fmt.Println(len(nums),cap(nums), nums[1])
	//nums = append(nums, 1)
	//fmt.Println(len(nums),cap(nums))
	//
	//nums = append(nums, 1)
	//fmt.Println(len(nums),cap(nums))


	//mapa := make(map[string]int, 2)
	//mapa["aaa"] = 1
	//mapa["bbb"] = 2
	//mapa["ccc"] = 2
	//fmt.Println(mapa, len(mapa))
	//a := twoPoint.RemoveElement(nums, 1)
	//fmt.Println(nums,a)

	//fmt.Println(math.Round(29/2))

	//prices := []int{7,1,5,3,6,4}
	//profit := dp.MaxProfit(prices)
	//fmt.Println(profit)

	//week08.Construct(5)
	//week08.Merge(1,2)
	//week08.Merge(2,3)
	//week08.Merge(3,4)
	//
	//week08.Find(1)
	//fmt.Println(week08.Fa)
	//strings.ToLower("HHH")
	//
	//name := "name1"
	//name2 := &name
	//
	//fmt.Println(*name2)

	//users := make([]*User, 5)
	//for _,s := range users {
		//s.Name = "aaa"

	//}

	//fmt.Println(users)

	//fmt.Println(week09.Violence("aaac", "ac"))
	//fmt.Println(week09.Kmp("aaac", "ac"))


	//for i,j := 1,0; i < m; i++ {
	//	for j > 0 && t[i] != t[j] {
	//		j = next[j - 1]
	//	}
	//	if t[i] == t[j] {
	//		j++
	//	}
	//	next[i] = j
	//}
	// agctagcagctagctg
	// 0123456789
	//t := "agctagcagctagctg"
	//m := len(t)
	//next := make([]int, m)
	//for i := 1; i < m; i ++ {
	//	k := next[i - 1]
	//	for t[i] != t[k] && k != 0 {
	//		k = next[k - 1]
	//	}
	//	if t[i] == t[k] {
	//		next[i] = k + 1
	//	} else {
	//		next[i] = 0
	//	}
	//}
	//
	//fmt.Println(next)


	//t := "fsdfsdfsdf"
	//m := len(t)
	//next := make([]int, m)
	//
	//for i := 1; i < m; i++ {
	//	k := next[i - 1]
	//	for t[k] != t[i] && k != 0 {
	//		k = next[k - 1]
	//	}
	//
	//	if t[k] == t[i] {
	//		next[i] = k + 1
	//	} else {
	//		next[i] = 0
	//	}
	//}
	//
	//fmt.Println(next)

	var n string
	var s string
	fmt.Scanln(&n)
	fmt.Scanln(&s)

	fmt.Println(n, "11111111")
	fmt.Println(s, "22222222222")
	arrS := strings.Split(s, " ")
	nums := make([]int, len(arrS))
	for i, v := range arrS {
		iv,_ := strconv.Atoi(v)
		nums[i] = iv
	}
	fmt.Println(nums)
	//fmt.Println(ironTower(8, []int{1 9 9 4 1 2 2 9}))
}

func ironTower(n int, nums []int) int {
	sum := make([]int, n + 1)
	sum[0] = 0

	for i := 1; i <= n; i++ {
		sum[i] = sum[i - 1] + nums[i - 1]
	}
	f := make([]int, n + 1)
	last := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if sum[i] - sum[j] >= last[j] {
				f[i] = f[j] + 1
				last[i] = sum[i] - sum[j]
			}
		}
	}
	return n - f[n]

}