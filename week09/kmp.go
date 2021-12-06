package week09

func kmp(s string, t string) int {
	if t == " " {
		return 0
	}

	n := len(s)
	m := len(t)

	next := make([]int, m)

	for i,j := 1,0; i < m; i++ {
		for j > 0 && t[i] != t[j] {
			j = next[j - 1]
		}
		if t[i] == t[j] {
			j++
		}
		next[i] = j
	}

	for i, j := 0,0; i < n; i++ {
		for j > 0 && t[j] != s[i] {
			j = next[j - 1]
		}

		if t[j] == s[i] {
			j++
		}

		if j == m {
			return i - m + 1
		}
	}

	return -1
}
