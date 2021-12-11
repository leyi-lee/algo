package week09

func Kmp(s string, t string) int {
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


func Violence(s string, t string) []int {
	n := len(s)
	m := len(t)

	ans := make([]int, 0)
	for i := 0; i < n - m + 1; i++ {
		w := i
		j := 0

		for j < m - 1 {
			if s[w] == t[j] {
				w++
				j++
			} else {
				break
			}
		}

		if s[w] == t[j] {
			ans = append(ans, i)
		}
	}

	return ans
}