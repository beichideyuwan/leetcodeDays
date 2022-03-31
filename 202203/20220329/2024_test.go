package _0220329

import "testing"

/*
一位老师正在出一场由 n道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F'表示）。
老师想增加学生对自己做出答案的不确定性，方法是最大化有 连续相同结果的题数。（也就是连续出现 true 或者连续出现 false）。
给你一个字符串answerKey，其中answerKey[i]是第 i个问题的正确结果。除此以外，还给你一个整数 k，表示你能进行以下操作的最多次数：
每次操作中，将问题的正确答案改为'T' 或者'F'（也就是将 answerKey[i] 改为'T'或者'F'）。
请你返回在不超过 k次操作的情况下，最大连续 'T'或者 'F'的数目。
示例 1：

输入：answerKey = "TTFF", k = 2
输出：4
解释：我们可以将两个 'F' 都变为 'T' ，得到 answerKey = "TTTT" 。
总共有四个连续的 'T' 。
示例 2：

输入：answerKey = "TFFT", k = 1
输出：3
解释：我们可以将最前面的 'T' 换成 'F' ，得到 answerKey = "FFFT" 。
或者，我们可以将第二个 'T' 换成 'F' ，得到 answerKey = "TFFF" 。
两种情况下，都有三个连续的 'F' 。
示例 3：

输入：answerKey = "TTFTTFTT", k = 1
输出：5
解释：我们可以将第一个 'F' 换成 'T' ，得到 answerKey = "TTTTTFTT" 。
或者我们可以将第二个 'F' 换成 'T' ，得到 answerKey = "TTFTTTTT" 。
两种情况下，都有五个连续的 'T' 。
*/
func TestMaxConsecutiveAnswers(t *testing.T) {
	answers := maxConsecutiveAnswers("TTFF", 2)
	t.Log(answers)
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	tCnt, fCnt, maxCnt, j, n := 0, 0, 0, 0, len(answerKey)
	for i := 0; i < n; i++ {
		if answerKey[i] == 'T' {
			tCnt++
			maxCnt = max(maxCnt, tCnt)
		} else {
			fCnt++
			maxCnt = max(maxCnt, fCnt)
		}
		if i-j+1 > maxCnt+k {
			if answerKey[j] == 'T' {
				tCnt--
			} else {
				fCnt--
			}
			j++
		}
	}
	return n - j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
