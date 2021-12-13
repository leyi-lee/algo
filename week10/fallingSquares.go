package week10

func fallingSquares(positions [][]int) []int {
	n := len(positions)
	qans := make([]int, n)

	for i := 0; i < n; i++ {
		left := positions[i][0]
		size := positions[i][1]
		right := left + size
		qans[i] += size

		for j := i+1; j < n; j++ {
			left2 := positions[j][0]
			size2 := positions[j][1]
			right2 := left2 + size2

			if left2 < right && left < right2 {
				qans[j] = max(qans[j], qans[i])
			}
		}
	}

	ans := make([]int, 0)
	cur := -1
	for _, x := range qans {
		cur := max(cur, x)
		ans = append(ans, cur)
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}