package stack

func Trap(heights []int) int { // 错误！！ 接雨水
	max := 0
	s := []rect{}
	for _, height := range heights{
		defaultWidth := 0
		for len(s) > 0 && s[len(s)-1].Height < height { // 单调递减
			defaultWidth += s[len(s)-1].Width
			top := s[len(s)-1]
			topHeight := top.Height
			s = s[:len(s)-1]

			up := topHeight
			if len(s) > 0 {
				second := s[len(s)-1]
				up = second.Height
			}

			rainHeight := up - topHeight
			if topHeight > up {
				rainHeight = topHeight - up
			}
			max += rainHeight * defaultWidth
		}
		s = append(s, rect{defaultWidth + 1, height})
	}

	return max
}
