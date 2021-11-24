package week08

var Fa []int

func Construct(n int) {
	Fa = make([]int, n)
	for i := 0; i < n; i++ {
		Fa[i] = i
	}
}

func Find(x int) int {
	if Fa[x] == x {
		return x
	}
	//return find(fa[x])
	// 合并路径 压缩路径
	Fa[x] = Find(Fa[x])
	return Fa[x]
}

func Merge(x int, y int) {
	Fa[Find(x)] = Find(y)
}