package _8

import (
	"testing"
)

/*
给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
示例 1：
输入：n = 13
输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
*/
func TestLexicalOrder(t *testing.T) {
	t.Log(lexicalOrder(13))
}

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i, _ := range ans {
		ans[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num+1 > n || num%10 == 9 {
				num /= 10
			}
			num++
		}
	}
	return ans
}
