package _02205

/*
基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序列start 默认是有效的，但是它并不一定会出现在基因库中。
*/
func diffOne(s, t string) (diff bool) {
	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return
}

func minMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}

	m := len(bank)
	adj := make([][]int, m)
	endIndex := -1
	for i, s := range bank {
		if s == end {
			endIndex = i
		}
		for j := i + 1; j < m; j++ {
			if diffOne(s, bank[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	if endIndex == -1 {
		return -1
	}

	var q []int
	vis := make([]bool, m)
	for i, s := range bank {
		if diffOne(start, s) {
			q = append(q, i)
			vis[i] = true
		}
	}
	for step := 1; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			if cur == endIndex {
				return step
			}
			for _, nxt := range adj[cur] {
				if !vis[nxt] {
					vis[nxt] = true
					q = append(q, nxt)
				}
			}
		}
	}
	return -1
}
