package week09

func toLowerCase(s string) string {
	ns := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			ns += string(s[i] + 'a' - 'A')
		} else {
			ns += string(s[i])
		}
	}
	return ns
}