package _0211213

import "math/rand"

/**
	冒泡
	插入
	归并
	快排
	选择
 */

// 冒泡 O(n^2)
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := 0; j < len(nums) - i - 1;j++ {
			if nums[j] > nums[j + 1] {
				temp := nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = temp
			}
		}
	}
	return nums
}

// 插入排序 最坏 O(n^2)  最好 O(n) 平均O( n * n )
func InsertSort(nums []int) []int {
	// 从 1 开始逐个往前比较
	for i := 1; i < len(nums); i++ {
		j := i
		for j > 0 && nums[j] < nums[j - 1] {
			nums[j], nums[j - 1] = nums[j - 1], nums[j]
			j--
		}
	}
	return nums
}

// 归并排序 nlogn
func MergeSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	MergeSort(nums, left, mid)
	MergeSort(nums, mid + 1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left int, mid int, right int)  {
	temp := make([]int, right - left + 1)

	i,j,k := left, mid + 1, 0
	for ; i <= mid && j <= right; k++ {
		if nums[i] < nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
	}

	// 补数据
	for ;i <= mid; k++ {
		temp[k] = nums[i]
		i++
	}

	for ; j <= right; i++ {
		temp[k] = nums[j]
		j++
	}

	for i := range temp{
		nums[left + i] = temp[i]
	}

}

// 快速排序
func QuickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := partition(nums, l, r)

	QuickSort(nums, l, pivot)
	QuickSort(nums, pivot + 1, r)
}

func partition(nums []int, l int, r int) int {
	pivot := rand.Intn(r - l + 1)
	pivotVal := nums[pivot]

	for l <= r {
		for nums[l] < pivotVal {
			l++
		}

		for nums[r] > pivotVal {
			r--
		}

		if l == r {
			break
		}

		if l < r {
			nums[l],nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	return r
}