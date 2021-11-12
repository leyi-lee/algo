package week06

/**
70. 爬楼梯
https://leetcode-cn.com/problems/climbing-stairs/description/
 */
// 递推公式  dp[i] = dp[i - 1] + dp[i - 2]
func climbStairs(n int) int {
	dp := make([]int, n + 1)

	dp[1] = 1
	if n <= 1 {
		return dp[n]
	}

	dp[2] = 2
	for i := 3;i <= n;i++ {
		dp[i] = dp[i - 2] + dp[i - 1]
	}


	return dp[n]
}