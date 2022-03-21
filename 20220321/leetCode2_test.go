package _0220321

import "testing"

/*
	看题解得出 有思路但不详细
给一个二进制字符串s。你可以按任意顺序执行以下两种操作任意次：

类型 1 ：删除 字符串s的第一个字符并将它 添加到字符串结尾。
类型 2 ：选择 字符串s中任意一个字符并将该字符反转，也就是如果值为'0'，则反转得到'1'，反之亦然。
请你返回使 s变成 交替 字符串的前提下，类型 2的 最少操作次数。

我们称一个字符串是 交替的，需要满足任意相邻字符都不同。

比方说，字符串"010" 和"1010"都是交替的，但是字符串"0100"不是。
*/
func TestMinFlips(t *testing.T) {
	flips := minFlips("0101111")
	t.Log(flips)
}

func minFlips(s string) int {
	n := len(s)
	ans := n
	// 枚举开头是 0 还是 1
	for head := byte('0'); head <= '1'; head++ {
		// 左边每个位置的不同字母个数
		leftDiff := make([]int, n)
		diff := 0
		for i := range s {
			if s[i] != head^byte(i&1) {
				diff++
			}
			leftDiff[i] = diff
		}

		// 右边每个位置的不同字母个数
		tail := head ^ 1
		diff = 0
		for i := n - 1; i >= 0; i-- {
			// 左边+右边即为整个字符串的不同字母个数，取最小值
			ans = min(ans, leftDiff[i]+diff)
			if s[i] != tail^byte((n-1-i)&1) {
				diff++
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
