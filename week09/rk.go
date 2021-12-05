package week09


type aaa interface {

}

func strStr(s, t string) int {
	if t == " " {
		return 0
	}

	n := len(s)
	m := len(t)

	s = " " + s
	t = " " + t

	p := int64(1e9 + 7)
	var tHash int64 = 0

	for i := 1; i <= m; i++ {
		tHash = (tHash * 131 + (int64(t[i]) - 'a' + 1)) % p
	}

	// 模板：预处理前缀hash
	sHash := make([]int64, n + 1)
	sHash[0] = 0

	p131 := make([]int64, n + 1)
	p131[0] = 1

	for i := 1; i <= n; i++ {
		sHash[i] = (sHash[i - 1] * 131 + (int64(s[i]) - 'a' + 1)) % p
		p131[i] = p131[i - 1] * 131 % p
	}

	for  i := m; i <= n; i++ {
		if calcHash(sHash, p131, p, i - m + 1, i) == tHash && s[i - m  + 1: i + 1] == t[1:] {
			return i - m
		}
	}
	return -1
}

func calcHash(H, p131 []int64, p int64, l, r int) int64 {
	return ((H[r] - H[l - 1] * p131[r - l + 1]) % p + p) % p
}