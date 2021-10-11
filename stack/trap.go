package stack

func Trap(heights []int) int { // 错误！！ 接雨水
	max := 0
	s := []rect{}
	for _, height := range heights{
		defaultWidth := 0
		for len(s) > 0 && s[len(s)-1].Height < height { // 单调递减
			defaultWidth += s[len(s)-1].Width // 累加宽度
			top := s[len(s)-1]
			topHeight := top.Height // 取出栈顶的高来
			s = s[:len(s)-1] // pop栈顶

			up := height
			if len(s) > 0 { // 说明未到头
				second := s[len(s)-1]
				up = second.Height
			} else {
				break
			}

			realHeight := Min(up, height)
			rainHeight := realHeight - topHeight
			max += rainHeight * defaultWidth
			//fmt.Println(realHeight, rainHeight, defaultWidth, max)
		}
		s = append(s, rect{defaultWidth + 1, height})
		//fmt.Println(s)
	}

	return max
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
/*
*
0,1,0,2,1,0,1,3,2,1,2,1
*/