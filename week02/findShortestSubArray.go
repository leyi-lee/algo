package week02

import "fmt"

/**
697. 数组的度
https://leetcode-cn.com/problems/degree-of-an-array/
 */
type positionCount struct {
	Count int
	Start int
	End int
}

/**
1、使用map 存储val出现的次数、首次出现的位置、最后出现的位置
2、找出count最大的值来，如果有相同的，使用end-start + 1比较谁最短
 */

func findShortestSubArray(nums []int) int {
	pMap := map[int]positionCount{}
	for i, val := range nums {
		if p,ok := pMap[val];ok {
			p.Count++
			p.End = i
			pMap[val] = p
			continue
		}

		pMap[val] = positionCount{1, i,i}
	}

	fmt.Println(pMap)

	ans := 0
	maxCount := 0
	for _, positionCount := range pMap {
		if positionCount.Count > maxCount {
			maxCount = positionCount.Count
			ans = positionCount.End - positionCount.Start + 1
			continue
		}

		if positionCount.Count == maxCount {
			if ans > positionCount.End - positionCount.Start + 1 {
				ans = positionCount.End - positionCount.Start + 1
			}
		}
	}

	return ans
}

