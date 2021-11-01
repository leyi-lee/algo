package week05

import "math"

/**
154. 寻找旋转排序数组中的最小值 II
https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/
 */

/**
	[4,5,6,1,2,3,4]
	[0,1,2,4,4,5,6]
按之前的逻辑写是满足的

小于等于
	[5,7,1,2,3,4,5,5]
     1,0,1,1,1,1,1,1
	不满足单调性
	[6,7,1,2,3,4,5,5,5]
如果还是小于最后一个
	0 0 1,1,1,1,0,0, 不满足单调性

	但是小于条件如果相等就让right自己在往左移动一下


这种是不满足的

 */
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := int(math.Floor(float64((left + right)/2)))
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else { // 如果相等就
			right = right - 1
		}
	}

	return nums[right]
}
