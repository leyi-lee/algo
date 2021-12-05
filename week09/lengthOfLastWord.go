package week09

func lengthOfLastWord(s string) int {
	n := len(s)
	right := n - 1
	left := n - 2

	for left >= 0 && right >= 0 && left < right{
		if s[right] == ' ' {
			right--
			left--
			continue
		}

		if s[left] == ' ' {
			break
		}

		left--
	}

	return right - left
}