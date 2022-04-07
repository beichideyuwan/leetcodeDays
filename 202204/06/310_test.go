package _6

/*
树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。
给你一棵包含n个节点的树，标记为0到n - 1 。给定数字n和一个有 n - 1 条无向边的 edges列表（每一个边都是一对标签）
，其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。
可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。
在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。
请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。
*/

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		deg[x]++
		deg[y]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}

	remainNodes := n
	for remainNodes > 2 {
		remainNodes -= len(q)
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, y := range g[x] {
				deg[y]--
				if deg[y] == 1 {
					q = append(q, y)
				}
			}
		}
	}
	return q
}
