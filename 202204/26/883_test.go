package _6

/*
在n x n的网格grid中，我们放置了一些与 x，y，z 三轴对齐的1 x 1 x 1立方体。
每个值v = grid[i][j]表示 v个正方体叠放在单元格(i, j)上。
现在，我们查看这些立方体在 xy、yz和 zx平面上的投影。
投影就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。
返回 所有三个投影的总面积 。
*/

func projectionArea(grid [][]int) int {
	var sumArea int
	for i, row := range grid {
		yzHeight, zxHeight := 0, 0
		for j, v := range row {
			if v > 0 {
				sumArea++
			}
			yzHeight = max(yzHeight, grid[j][i])
			zxHeight = max(zxHeight, v)
		}
		sumArea += yzHeight
		sumArea += zxHeight
	}
	return sumArea
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
