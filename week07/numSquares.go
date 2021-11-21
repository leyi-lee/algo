package week07

func numSquares(n int) int {
	f := make([]int, n + 1)
	for i := range f {
		f[i] = 10000000
	}

	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j * j <= i;j++ {
			f[i] = min(f[i], f[i - j * j] + 1)
		}
	}

	return f[n]
}

func min(x int ,y int) int {
	if x > y {
		return y
	}
	return x
}