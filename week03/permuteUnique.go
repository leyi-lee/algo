package week03

import "sort"

/**
全排列
https://leetcode-cn.com/problems/permutations-ii/
 */

var ans [][]int
var used []bool
var chosen []int


func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 先排序把相同的值放一起
	for i := 0; i < len(nums);i++ {
		used = append(used, false)
	}

	ans = [][]int{}
	chosen = []int{}
	recur(nums, 0)
	return ans
}

func recur(nums []int, pos int) {
	// 返回答案
	if pos == len(nums) {
		p := make([]int, len(chosen))
		copy(p, chosen)
		ans = append(ans, p)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 前一个数未用且相等也跳过去
		if used[i] || i > 0 && !used[i - 1] && nums[i] == nums[i - 1] {
			continue
		}

		chosen = append(chosen, nums[i])
		used[i] = true
		recur(nums, pos + 1)

		used[i] = false
		chosen = chosen[:len(chosen) - 1]
	}
}