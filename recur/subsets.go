package recur

var chosen []int
var ans [][]int

func Subsets(nums []int) [][]int {
	chosen = []int{}
	ans = [][]int{}

	recur(nums, 0)
	return ans
}

func recur(nums []int, pos int) {
	if pos == len(nums) {
		p := make([]int, len(chosen))
		copy(p, chosen)
		ans = append(ans, p)
		return
	}

	chosen = append(chosen, nums[pos])
	recur(nums, pos + 1)
	chosen = chosen[:len(chosen) - 1]

	recur(nums, pos+1)
}

