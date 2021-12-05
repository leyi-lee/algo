package week09

func firstUniqChar(s string) int {
	sHash := map[string]int{}

	for _,v := range s {
		if _,ok := sHash[string(v)]; ok {
			sHash[string(v)] += 1
		} else {
			sHash[string(v)] = 1
		}
	}

	for i,v := range s {
		if _,ok := sHash[string(v)]; ok && sHash[string(v)] == 1 {
			return i
		}
	}

	return -1
}