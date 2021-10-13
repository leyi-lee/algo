package PrefixSum

/**
53
https://leetcode-cn.com/problems/maximum-subarray/
 */

func MaxSubArray(nums []int) int {

	s := []int{0}
	preMin := []int{}
	max := 0


	for i := 1; i <= len(nums); i++ {
		s = append(s, s[i-1] + nums[i-1])
	}

	for i := 1; i < len(nums); i++ {
		preMin = append(preMin, min(preMin[i-1], s[i]))
	}

	for i := 1; i < len(nums); i++ {
		max = Max(max, s[i] - preMin[i-1])
	}
	return max
}

func min(x int, y int) int {
	if x > y {
		return y
	}

	return x
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}