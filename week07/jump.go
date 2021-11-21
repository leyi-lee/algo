package week07

/**
45. 跳跃游戏 II
https://leetcode-cn.com/problems/jump-game-ii/
 */

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j + nums[j] >= i {
				dp[i] = min(dp[i], dp[j] + 1)
			}
		}
	}
	fmt.Println(dp)
	return dp[n - 1]
}