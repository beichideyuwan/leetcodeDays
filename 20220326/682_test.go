package _0220326

import (
	"strconv"
	"testing"
)

/*
你现在是一场采用特殊赛制棒球比赛的记录员。这场比赛由若干回合组成，过去几回合的得分可能会影响以后几回合的得分。

比赛开始时，记录是空白的。你会得到一个记录操作的字符串列表 ops，其中 ops[i] 是你需要记录的第 i 项操作，ops 遵循下述规则：

整数 x - 表示本回合新获得分数 x
"+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
"D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
"C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。
请你返回记录中所有得分的总和。
*/
func TestCalPoints(t *testing.T) {
	points := calPoints([]string{"5", "2", "C", "D", "+"})
	t.Log(points)
}
/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了78.38%的用户
*/
func calPoints(ops []string) int {
	s:=make([]int,0)
	for _, op := range ops {
		switch op {
		case "+":
			s = append(s,s[len(s)-1]+s[len(s)-2])
		case "D":
			s = append(s,s[len(s)-1]*2)
		case "C":
			s = s[:len(s)-1]
		default:
			itoa,_ := strconv.Atoi(op)
			s = append(s, itoa)
		}
	}
	var ans int
	for _, i := range s {
		ans +=i
	}
	return ans
}
