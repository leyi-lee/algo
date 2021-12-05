package week09

func numJewelsInStones(jewels string, stones string) int {
	jHash := map[string]int{}

	for _,v := range jewels {
		jHash[string(v)] = 1
	}

	ans := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := jHash[string(stones[i])];ok {
			ans++
		}
	}
	return ans
}