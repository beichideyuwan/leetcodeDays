package _02204

/*
给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。
你需要返回能表示矩阵的 四叉树 的根结点。
注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。
四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：
val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False；
isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。
*/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	var dfs func([][]int, int, int) *Node
	dfs = func(rows [][]int, c0, c1 int) *Node {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] { // 不是叶节点
					rMid, cMid := len(rows)/2, (c0+c1)/2
					return &Node{
						true,
						false,
						dfs(rows[:rMid], c0, cMid),
						dfs(rows[:rMid], cMid, c1),
						dfs(rows[rMid:], c0, cMid),
						dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		// 是叶节点
		return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
	}
	return dfs(grid, 0, len(grid))
}
