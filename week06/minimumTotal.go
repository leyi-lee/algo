package week06

/**
120. 三角形最小路径和
https://leetcode-cn.com/problems/triangle/description/
 */
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0;i < n;i++ {
		f[i] = make([]int, n)
	}

	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i][0] = f[i - 1][0] + triangle[i][0] // 两个边直接按前一个加
		for j := 1; j < i;j++ {
			f[i][j] = min(f[i - 1][j], f[i - 1][j - 1]) + triangle[i][j]
		}
		f[i][i] = f[i - 1][i - 1] + triangle[i][i] // 两个边直接按前一个加
	}


	ans := 10000
	for i := 0; i < n; i++ {
		ans = min(f[n - 1][i], ans)
	}

	return ans
}

func min(x int, y int) int {
	if x > y {
		return y
	}

	return x
}