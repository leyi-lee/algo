package week01

import "strconv"

/**
85
https://leetcode-cn.com/problems/maximal-rectangle/
 */

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
			if row == 0 {
				colInt, _ := strconv.Atoi(string(matrix[row][col]))
				heights = append(heights, colInt)
				continue
			}
			//isNotAppend := len(heights) != 0 && len(heights)-1 >= col // 判断底是否有了，其实就是遍历的第一行

			if string(matrix[row][col]) == "1" { // 如果为1 在基础上加一
				heights[col] += 1
			} else {
				heights[col] = 0
			}
		}
		newRect := largestRectArea(heights)
		if newRect > max {
			max = newRect
		}
	}
	return max
}

// 重写 柱形图最大矩形
type rect struct {
	Width int
	Height int
}

func largestRectArea(heights []int) int {
	max := 0
	s := []rect{}
	heights = append(heights, 0)
	for _, height := range heights {
		defaultWidth := 0
		if len(s) > 0 && s[len(s)-1].Height >= height { // 栈顶大于后插入的柱子高度开始计算
			defaultWidth += s[len(s) - 1].Width
			newRect := defaultWidth * s[len(s) - 1].Height
			if newRect > max {
				max = newRect
			}
			s = s[:len(s) -1]
		}

		s = append(s, rect{defaultWidth + 1, height})
	}

	return max
}