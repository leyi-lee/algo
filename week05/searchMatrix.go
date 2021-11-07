package week05

import "math"

func searchMatrix(matrix [][]int, target int) bool {
	result := false
	for _,row := range matrix {
		if target >= row[0] && row[len(row) - 1] >= target {
			result = binarySearch(row, target)
			break
		}
	}
	return result
}

func binarySearch(nums []int, target int) bool {
	left := 0
	right := len(nums)

	for left < right {
		mid := int(math.Floor(float64((left + right)/2)))
		if nums[mid] == target {
			return true
		}

		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}