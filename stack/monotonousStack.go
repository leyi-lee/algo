package stack

import "fmt"

type rect struct {
	Width int
	Height int
}

// 柱状图中最大矩形
func LargestRectangleArea(heights []int) int {
	max := 0
	s := []rect{} // 栈存放矩形
	heights = append(heights, 0)
	for _, height := range heights {
		defaultWidth := 0
		for len(s) > 0 && s[len(s)-1].Height > height {
			defaultWidth += s[len(s)-1].Width
			newRect := defaultWidth * s[len(s)-1].Height // 出错！！！计算的是栈顶的高度
			if newRect >= max {
				max = newRect
			}
			s = s[:len(s)-1]
		}

		// 出错 ！！！ defaultWidth + 1  为什么这要加一，就是你要开始计算或者合并柱状图时，最后入栈的要加上破坏单调递增的那个宽度
		s = append(s, rect{Width: defaultWidth+1, Height: height})
		fmt.Println(s)
	}

	return max
}

/*
[]int{2,1,5,6,2,3}
s的结构
[{1 2}] 第一次先放入第一位
[{2 1}] 第二次 1比2 小，开始计算 宽是2 高是1
[{2 1} {1 5}] 第三次 递增 继续放
[{2 1} {1 5} {1 6}] 第四次递增 继续放
[{2 1} {3 2}] 第五次 2 比 栈顶的6小，开始计算计算到第一个入栈的柱子 1 比2小，计算了三个柱子所以分别是5，6，2所以宽为3 高为2 入栈
[{2 1} {3 2} {1 3}] 第七次 递增入栈
[{7 0}] 0 无高度，计算前面的
*/

