package week08

/**
684. 冗余连接
https://leetcode-cn.com/problems/redundant-connection/
 */

/**
并查集
 */

var fa []int

func findRedundantConnection(edges [][]int) []int {
	fa = make([]int, len(edges) + 1)

	for i := range fa {
		fa[i] = i
	}

	for _,item := range edges {
		if !merge(item[0], item[1]) {
			return item
		}
	}

	return []int{}
}

func find(x int) int {
	if fa[x] == x {
		return x
	}

	fa[x] = find(fa[x])

	return fa[x]
}

func merge(x int, y int) bool {
	if find(x) == find(y) {
		return false
	}

	fa[find(x)] = find(y)
	return true
}