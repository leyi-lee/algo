package week08

import "fmt"

/**
200. 岛屿数量
https://leetcode-cn.com/problems/number-of-islands/
 */

var fat []int
var m int
var n int

func numIslands(grid [][]byte) int {
	m = len(grid)
	n = len(grid[0])
	fat = make([]int, m * n)
	for i := 0;i < m * n; i++ {
		fat[i] = i
	}

	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	for i := 0 ; i < m; i++ {
		for j :=0; j < n; j++ {
			if string(grid[i][j]) == "0" {
				continue
			}

			for k := 0; k < 4; k++ {
				ni := i + dx[k]
				nj := j + dy[k]

				if ni < 0 || nj < 0 || ni >= m || nj >= n  {
					continue
				}

				if string(grid[ni][nj]) == "0" {
					continue
				}

				union(node(i, j), node(ni, nj)) // 链接
			}
		}
	}
	fmt.Println(fat)

	ans := 0
	for i := 0 ; i < m; i++ {
		for j :=0; j < n; j++ {
			// fmt.Println(node(i, j))
			findRoot := finds(node(i, j))
			fmt.Println(node(i, j),findRoot)
			if findRoot == node(i,j) && string(grid[i][j]) != "0" { // 判断根是自己，且节点不为0
				ans++
			}
		}
	}

	return ans
}

func node(i int, j int) int {
	return i * n + j
}

func finds(x int) int {
	if fat[x] == x {
		return x
	}
	//return find(fa[x])
	// 合并路径 压缩路径
	fat[x] = finds(fat[x])
	return fat[x]
}

func union(x int, y int) {
	x = finds(x)
	y = finds(y)
	if x != y {
		fat[x] = y
	}
}