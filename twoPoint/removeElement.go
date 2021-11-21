package twoPoint

import (
	"fmt"
	"sort"
)

func RemoveElement(nums []int, target int) int {
	slowIndex := 0
	for f := 0;f < len(nums);f++ {
		if nums[f] != target {
			nums[slowIndex] = nums[f]
			slowIndex++
		}
	}
	return slowIndex
}


func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			if nums[i] + nums[left] + nums[right] > 0 {
				right--
			} else if nums[i] + nums[left] + nums[right] < 0 {
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				for right > left && nums[right] == nums[right - 1] {
					right--
				}

				for right > left && nums[left] == nums[left - 1] {
					left++
				}

				right--
				left++
			}
		}
	}
	return ans
}

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans :=[][]int{}
	for k := 0; k < len(nums);k++ {

		for i := k+1;i < len(nums);i++ {

			left := i+1
			right := len(nums) - 1

			for left < right {
				if nums[k] + nums[i] + nums[left] + nums[right] > target {
					right--
				} else if nums[k] + nums[i] + nums[left] + nums[right] < target {
					left++
				} else {
					ans = append(ans, []int{nums[k],nums[i],nums[left],nums[right]})

					for right > left && nums[right] == nums[right - 1] {
						right--
					}

					for right > left && nums[left] == nums[left + 1] {
						left++
					}

					right--
					left++
				}
			}
		}
	}
	return ans
}

func SortArrRemove() {
	nums := []int{1,1,1,2,2,3,4,4,5,6}
	sort.Ints(nums)
	n := len(nums)

	k := 0
	m := 0
	for i := k + 1; i < n; i++ {
		if k - m >= 0 && nums[k - m] == nums[i] {
			continue
		}

		nums[k + 1] = nums[i]
		k++
	}

	fmt.Println(nums[:k + 1])
}