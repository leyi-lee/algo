package week02

import "fmt"

/**
1074. 元素和为目标值的子矩阵数量
https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/
 */

/**
解题思路
分别按行和列统计 和 为k的子数组数量 错误思想！！！
直接使用 https://leetcode-cn.com/problems/subarray-sum-equals-k/
 */

func NumSubmatrixSumTarget(matrix [][]int, target int) int {
	ans := 0
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			fmt.Println( matrix[i:])
			for c, v := range row {
				sum[c] += v
			}
			//fmt.Println(sum, 1)
			ans += subArraySum(sum, target)
		}
	}
	return ans
}

func subArraySum(nums []int, k int) int {
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