package dsort

import "math/rand"

func bubbleSort(nums []int) []int {
	n := len(nums) - 1
	for i := 0;i < n - 1;i++ {
		for j := 0; j < n - 1 - i;j++{
			if nums[j] > nums[j+1] {
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
		}
	}
	return nums
}

// 可用希尔排序优化
/**
每走一步都往前面比较替换直到满足条件
 */
func InsertSort(nums []int) []int {
	for i := 1; i < len(nums);i++ {
		j := i
		for j > 0 && nums[j] < nums[j-1] {
			nums[j],nums[j-1] = nums[j-1], nums[j]
			j--
		}
	}
	return nums
}

// 可用小堆排序优化
func SelectSort(nums []int) []int {
	for i := 0;i < len(nums) - 1;i++{

		minIndex := i
		for j := i + 1;j < len(nums);j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}


		if i != minIndex {
			nums[i],nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}

// 归并排序
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

	i, j ,k := left, mid + 1, 0

	// 合并两个有序数组
	for ;i <= mid && j <= right;k++ {
		if nums[i] < nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
	}

	// 补数据
	for ;i <= mid;k++{
		temp[k] = nums[i]
		i++
	}

	for ;j <= right;k++{
		temp[k] = nums[j]
		j++
	}

	for i := range temp {
		nums[left + i] = temp[i]
	}
}

// 快排
func quickSort(nums []int, l int, r int)  {
	if l >= r {
		return
	}

	pivot := partition(nums, l, r)
	quickSort(nums, l, pivot)
	quickSort(nums,pivot + 1, r)
}

func partition(nums []int, l int, r int) int {
	pivot := l + rand.Intn(r - l + 1)
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
			nums[l],nums[r] = nums[r],nums[l]
			l++
			r--
		}
	}

	return r
}