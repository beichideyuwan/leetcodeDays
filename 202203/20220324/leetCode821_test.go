package _0220324

/*
821. 字符的最短距离
给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。

返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。

case
两个下标i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数
输入：s = "loveleetcode", c = "e"
输出：[3,2,1,0,1,0,0,1,2,2,1,0]
解释：字符 'e' 出现在下标 3、5、6 和 11 处（下标从 0 开始计数）。
距下标 0 最近的 'e' 出现在下标 3 ，所以距离为 abs(0 - 3) = 3 。
距下标 1 最近的 'e' 出现在下标 3 ，所以距离为 abs(1 - 3) = 2 。
对于下标 4 ，出现在下标 3 和下标 5 处的 'e' 都离它最近，但距离是一样的 abs(4 - 3) == abs(4 - 5) = 1 。
距下标 8 最近的 'e' 出现在下标 6 ，所以距离为 abs(8 - 6) = 2 。
*/

func shortestToChar(S string, C byte) []int {
	ret := make([]int, len(S))
	for index := 0; index < len(S); index++ {
		// 找到C的位置
		if S[index] == C {
			ret[index] = 0

			// 从C的位置向左遍历,直到到达数组最左或者到达下一个C
			for m := index - 1; m >= 0; m-- {
				if S[m] == C {
					break
				} else {
					dd := ret[m+1] + 1
					if ret[m] == 0 || dd < ret[m] {
						ret[m] = dd
					}
				}
			}

			// 从C的位置向右遍历,直到到达数组最右或者到达下一个C,
			for n := index + 1; n < len(S); n++ {
				if S[n] == C {
					break
				} else {
					dd := ret[n-1] + 1
					if ret[n] == 0 || dd < ret[n] {
						ret[n] = dd
					}
				}
			}
		}
	}
	return ret
}