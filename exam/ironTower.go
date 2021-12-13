package exam
// 铁塔

func ironTower(n int, nums []int) int {
	sum := make([]int, n + 1)
	sum[0] = 0

	for i := 1; i <= n; i++ {
		sum[i] = sum[i - 1] + nums[i - 1]
	}
	f := make([]int, n + 1)
	last := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if sum[i] - sum[j] >= last[j] {
				f[i] = f[j] + 1
				last[i] = sum[i] - sum[j]
			}
		}
	}
	return n - f[n]

}