package hash

/**
leetCode 1
https://leetcode-cn.com/problems/two-sum/description/
 */

func TwoSum(nums []int, target int) []int {
	unOrderSet := map[int]int{}

	for i, num := range nums {
		diffNum := target - num
		if _, ok := unOrderSet[diffNum]; ok {
			hitIndex,_ := unOrderSet[diffNum]
			return []int{hitIndex, i}
		}

		unOrderSet[num] = i
	}
	return []int{}
}
