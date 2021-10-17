package week02

/**
560. 和为 K 的子数组
https://leetcode-cn.com/problems/subarray-sum-equals-k/
 */

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums) + 1)
	sums[0] = 0

	for i := 1;i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	ans := 0
	mapSum := map[int]int{}
	for i := 0;i <= len(nums); i++ {
		if _, ok := mapSum[sums[i] - k];ok {
			ans += mapSum[sums[i] - k]
		}
		mapSum[sums[i]] += 1
	}
	return ans
}
