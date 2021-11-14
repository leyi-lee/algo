package week06

import "fmt"

/**

!!! 有问题
673. 最长递增子序列的个数
https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
 */
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	countF := make([]int, n)
	maxInt := 0

	for i := 0; i < n; i++ {
		f[i] = 1
		countF[i] = 1
	}

	for i := 1; i < n; i++ {
		maxIv := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j] + 1)

				if maxIv < f[i] {
					countF[i] = 1
					maxIv = f[i]
				} else {
					countF[i]++
				}
			} else if nums[i] == nums[j] {
				countF[i]++
			}
		}

		maxInt = max(maxInt, f[i]) // 记录最长长度
	}

	fmt.Println(countF)
	fmt.Println(f)
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, countF[i])
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}