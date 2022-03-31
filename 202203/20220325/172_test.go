package _0220325

import "testing"

/*
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
*/
func TestTrailingZeroes(t *testing.T) {
	t.Log(trailingZeroes(30))
}
func trailingZeroes(n int) (s int) {
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			s++
		}
	}
	return
}