package week05

import (
	"math"
)

/**
875. 爱吃香蕉的珂珂
https://leetcode-cn.com/problems/koko-eating-bananas/
 */

func MinEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1
	for _,pile := range piles {
		right = max(right, pile)
	}

	for left < right {
		mid := int(math.Floor(float64((left + right)/2)))
		if validateK(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func validateK(piles []int, h int, k int) bool {
	allH := 0
	for _,pile := range piles {
		sH := 0
		if pile % k == 0 {// !!! 问题 向上取整
			sH = pile / k
		} else {
			sH = int(math.Floor(float64(pile/k) + 1))
		}
		allH += sH
	}
	return allH <= h
}