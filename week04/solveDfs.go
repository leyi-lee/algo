package week04

//var m int
//var n int
//var visited [][]bool
//var ans [][]int
//var q [][]int
//var isNeed [][]int
var isOk bool

func solveDfs(board [][]byte) {
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
			isOk = true
			if string(board[i][j]) == "O" && !visited[i][j] {
				dfs(i, j, board)
			}

			if len(isNeed) > 0 && isOk {
				ans = append(ans, isNeed...)
			}
		}
	}

	var x byte = 'X'
	for _, rc := range ans {
		board[rc[0]][rc[1]] = x
	}

}

func dfs(rx int, ry int, board [][]byte)  {
	if rx < 0 || rx >= m || ry < 0 || ry >= n || visited[rx][ry] || string(board[rx][ry]) == "X" {
		return
	}

	if (rx == 0 || rx == m - 1 || ry == 0 || ry == n - 1) && string(board[rx][ry]) == "O" {
		isOk = false
	}

	visited[rx][ry] = true
	isNeed = append(isNeed, []int{rx, ry})

	dx := []int{-1, 0 ,0, 1}
	dy := []int{0, -1 ,1, 0}

	for i := 0; i < 4;i++ {
		nx := rx + dx[i]
		ny := ry + dy[i]
		dfs(nx, ny, board)
	}
}