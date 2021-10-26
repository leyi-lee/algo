package week04

/**
130. 被围绕的区域
https://leetcode-cn.com/problems/surrounded-regions/
 */

var m int
var n int
var visited [][]bool
var ans [][]int
var q [][]int
var isNeed [][]int

/**
 主要思想找O
	不从边找 行 i 从 1 开始 且 小于 m - 1 列同理 第一行和最后一行，第一列和最后一列
	开始从一个为O的点找，bfs 从队列取头取出来开始走，走到边不停止继续让走完，看看区域有多大，如果碰到边，记录一下
	最后从走过的里面判断没走过边的 放入答案
 */

func solve(board [][]byte) {
	m = len(board)
	n = len(board[0])
	visited = make([][]bool, m)
	ans = [][]int{}
	isNeed = [][]int{}
	for i := 0; i < m; i++ { // 初始化访问数据
		itemInit := make([]bool, n)
		for j := 0; j < n; j++ {
			itemInit[j] = false
		}
		visited[i] = itemInit
	}

	for i := 1; i < m - 1 ; i++ {
		for j := 1; j < n - 1; j++ {
			isNeed = [][]int{}
			if string(board[i][j]) == "O" && !visited[i][j] {
				bfs(i, j, board)
			}
		}
	}

	var x byte = 'X'
	for _, rc := range ans {
		board[rc[0]][rc[1]] = x
	}

}

func bfs(nx int, ny int, board [][]byte)  {
	visited[nx][ny] = true
	q = append(q, []int{nx, ny})
	isNeed = append(q, []int{nx, ny})

	dx := []int{-1, 0 ,0, 1}
	dy := []int{0, -1 ,1, 0}

	isOk := true
	for len(q) > 0 {
		head := q[0]
		q = q[1:]

		for i := 0; i < 4;i++ {
			nx = head[0] + dx[i]
			ny = head[1] + dy[i]
			if nx < 0 || nx >= m || ny < 0 || ny >= n || visited[nx][ny] {
				continue
			}

			if string(board[nx][ny]) == "X" {
				continue
			}

			if (nx == 0 || nx == m - 1 || ny == 0 || ny == n - 1) && string(board[nx][ny]) == "O" {
				isOk = false
			}

			visited[nx][ny] = true
			q = append(q, []int{nx, ny})
			isNeed = append(isNeed, []int{nx, ny})
		}
	}

	if len(isNeed) > 0 && isOk {
		ans = append(ans, isNeed...)
	}
}