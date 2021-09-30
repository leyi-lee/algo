package week01

func PlusOne(digits []int) []int {
	for i := len(digits)-1;i >= 0; i-- {
		digits[i]++
		if digits[i] % 10 != 0 { // 只要不等于0 直接停掉，不进一
			return digits
		}

		digits[i] = 0 // 如果执行到这说明这个位是满10的直接等于0
	}

	return append([]int{1},digits...) // 在这一步返回说明最后最高位也是满10的在头部位置放1
}