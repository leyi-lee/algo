package week01

import (
	"fmt"
	"myAlgo/stack"
)

// 主要思想单调栈求最大值
func MaximalRectangle(matrix [][]byte) int {
	max := 0
	mLen := len(matrix) // 如果为0
	if mLen == 0 {
		return max
	}

	itemLen := len(matrix[0])
	heights := []int{}
	for row := 0; row < mLen; row++ { // 遍历数量直接取值（不直接遍历matrix）
		for col := 0; col < itemLen; col++ {
			isNotAppend := len(heights) != 0 && len(heights)-1 >= col

			if string(matrix[row][col]) == "1" {
				if isNotAppend {
					heights[col] += 1
				} else {
					heights = append(heights, 1)
				}
			} else {
				if isNotAppend {
					heights[col] = 0
				} else {
					heights = append(heights, 0)
				}
			}
		}
		newRect := stack.LargestRectangleArea(heights)
		fmt.Println(heights, newRect)
		if newRect > max {
			max = newRect
		}
	}
	return max
}