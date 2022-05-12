package _02205

/*
由范围 [0,n] 内所有整数组成的 n + 1 个整数的排列序列可以表示为长度为 n 的字符串 s ，其中:
如果perm[i] < perm[i + 1]，那么s[i] == 'I'
如果perm[i] > perm[i + 1]，那么 s[i] == 'D'
给定一个字符串 s ，重构排列perm 并返回它。如果有多个有效排列perm，则返回其中 任何一个 。
*/
func diStringMatch(s string) []int {
	n := len(s)
	perm := make([]int, n+1)
	lo, hi := 0, n
	for i, ch := range s {
		if ch == 'I' {
			perm[i] = lo
			lo++
		} else {
			perm[i] = hi
			hi--
		}
	}
	perm[n] = lo // 最后剩下一个数，此时 lo == hi
	return perm
}
