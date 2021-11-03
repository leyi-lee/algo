package week05

import "math"

/**
1011. 在 D 天内送达包裹的能力
https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
 */

/**
问题转换
	验证在哪些运载重量下，可以使包裹在days天运输完
	范围取值left最大重量，right所有重量
	满足单调性使用二分答案
	在满足条件的情况下找最小的
 */

func shipWithinDays(weights []int, days int) int {
	left := 0
	right := 0

	for _,weight := range weights {
		left = max(left, weight)
		right += weight
	}

	for left < right {
		mid := int(math.Floor(float64((left + right)/2)))
		if validateWeights(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func validateWeights(weights []int, days int, vWeight int) bool {
	allowDay := 0
	allWeight := 0

	for _,weight := range weights {
		if allWeight + weight > vWeight {
			allowDay++
			allWeight = weight
		} else {
			allWeight += weight
		}
	}

	return allowDay < days
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}