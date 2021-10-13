package queue

/**
LeetCode 239
https://leetcode-cn.com/problems/sliding-window-maximum/
 */

func MaxSlidingWindow(nums []int, k int) []int {

	d := []int{}
	ans := []int{}
	for i := 0; i < len(nums) ; i++ {

		for len(d) > 0 && d[0] <= i - k {
			d = d[1:]
		}

		for len(d) > 0 && nums[d[len(d)-1]] <= nums[i] {
			d = d[:len(d)-1]
		}

		d = append(d, i)

		if k - 1 <= i {
			ans = append(ans, nums[d[0]])
		}
	}
	return ans
}

/*
1,3,-1,-3,5,3,6,7
     1 0 1 0 1 0
 0 1 2 3 4 5 6 7


正常入单调队列下标

队尾取出来与新入的值作比较 如果后面的大 则把队尾pop
然后插入新的i

如果 i已经大于等于 k - 1,为什么是k-1，因为一次滑动1步骤，只要大于这个数值的i都需要计算一次
换句话说得取队列的头部了


判断队头是否已经超过了分组的边界，如果判断，就是 队头的下标 要比当前的i-k还要小
比如说 i= 0 1 2 3 4，假如k是3，i = 4了那么如果队头存的0或者1就都没用了


// 原题不变，但是每次滑动是n

*/

func MaxSlidingWindowSlide(nums []int, k int, step int) []int {
	d := []int{}
	ans  := []int{}

	for i := 0; i < len(nums); i++ {
		for len(d) > 0 && d[0] <= i - k  {
			d = d[1:]
		}

		for len(d) > 0 && nums[d[len(d)-1]] <= nums[i] {
			d = d[:len(d)-1]
		}

		d = append(d, i)

		if i >= k - 1 {
			ans = append(ans, nums[d[0]])
		}
	}

	return ans
}