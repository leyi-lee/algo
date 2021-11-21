package week07

/**
55. 跳跃游戏
https://leetcode-cn.com/problems/jump-game/
 */

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)

	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j + nums[j] >= i {
				fmt.Println(dp[j] + 1)
				dp[i] = min(dp[i], dp[j] + 1)
			}
		}
	}

	return dp[n - 1] != n + 1
}