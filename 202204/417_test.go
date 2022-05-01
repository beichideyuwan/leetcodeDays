package _02204

/*
有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。“太平洋”处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵heights，heights[r][c]表示坐标
(r, c) 上单元格 高于海平面的高度 。
岛上雨水较多，如果相邻单元格的高度
小于或等于 当前单元格的高度，
雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
返回网格坐标 result的 2D 列表 ，其中result[i] = [ri, ci]表示雨水从单元格 (ri, ci) 流动 既可流向太平洋也可流向大西洋
*/
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, d := range dirs {
			if nx, ny := x+d.x, y+d.y; 0 <= nx && nx < m && 0 <= ny && ny < n && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 1; j < n; j++ {
		dfs(0, j, pacific)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n-1; j++ {
		dfs(m-1, j, atlantic)
	}

	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
