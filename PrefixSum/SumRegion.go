package PrefixSum

/**
304. 二维区域和检索 - 矩阵不可变
https://leetcode-cn.com/problems/range-sum-query-2d-immutable/
 */

type NumMatrix struct {
	Sums [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))

	for i, row := range matrix {
		sums[i] = make([]int, len(row) + 1)

		for j := 0;j < len(row);j++ {
			sums[i][j+1] = sums[i][j] + row[j]
		}
	}

	return NumMatrix{sums}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2;i++ {
		sum += this.Sums[i][col2+1] - this.Sums[i][col1]
	}
	return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */