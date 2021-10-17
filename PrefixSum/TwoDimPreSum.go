package PrefixSum

func TwoDimPreSum(nums [][]int) [][]int {

	rows := len(nums)
	cols := len(nums[0])

	sums := make([][]int, rows + 1)
	for i := range sums {
		sums[i] = make([]int, cols + 1)
	}
	sums[0][0] = 0
	for i :=1; i <= rows; i++ {
		for j :=1; j <= cols; j++ {
			sums[i][j] = nums[i-1][j-1] + sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1]
		}
	}

	return sums
}

func PreSum(nums []int) []int {
	sums := make([]int, len(nums) + 1)
	sums[0] = 0

	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	return sums
}

func PreDiff(nums []int) []int {
	diffs := make([]int, len(nums))
	diffs[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		diffs[i] = nums[i] - nums[i-1]
	}

	return diffs
}


func TwoDimPreDiff(nums [][]int) [][]int {

	rows := len(nums)
	cols := len(nums[0])

	diffs := make([][]int, rows + 1)
	for i := range diffs {
		diffs[i] = make([]int, cols + 1)
	}
	diffs[0][0] = 0
	for i :=1; i <= rows; i++ {
		for j :=1; j <= cols; j++ {
			diffs[i][j] = nums[i-1][j-1] - nums[i-1][j] - nums[i][j-1] + nums[i-1][j-1]
		}
	}

	return diffs
}