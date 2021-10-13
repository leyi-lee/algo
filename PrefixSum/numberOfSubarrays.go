package PrefixSum

func NumberOfSubArrays(nums []int, k int) int {
	ans := 0

	// 1,1,2,1,1
	sums := []int{0} // 前缀和
	for i := 1; i <= len(nums); i++ {
		sums = append(sums,sums[i-1] + (nums[i-1] % 2))
	}

	count := map[int]int{}
	for i := 0;i <= len(nums);i++ { // 统计结果出现的次数
		count[sums[i]]++
	}

	for i := 1; i <= len(nums); i++ {
		if sums[i] - k >= 0 {
			ans += count[sums[i] - k] // 累加满足s[j] 的统计数据
		}
	}
	return ans
}
