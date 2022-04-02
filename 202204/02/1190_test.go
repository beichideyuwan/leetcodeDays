package _2

import "testing"

/*
1190. 反转每对括号间的子串
给出一个字符串s（仅含有小写英文字母和括号）。
请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
注意，您的结果中 不应 包含任何括号。
示例 1：
输入：s = "(abcd)"
输出："dcba"
示例 2：
输入：s = "(u(love)i)"
输出："iloveu"
解释：先反转子字符串 "love" ，然后反转整个字符串。
示例 3：
输入：s = "(ed(et(oc))el)"
输出："leetcode"
解释：先反转子字符串 "oc" ，接着反转 "etco" ，然后反转整个字符串。
示例 4：
输入：s = "a(bcdefghijkl(mno)p)q"
输出："apmnolkjihgfedcbq"
*/

func TestReverseParentheses(t *testing.T) {
	t.Log(reverseParentheses("(ed(et(oc))el)"))
}

func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stack := []int{}
	for i, b := range s {
		if b == '(' {
			stack = append(stack, i)
		} else if b == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}

	ans := []byte{}
	for i, step := 0, 1; i < n; i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			step = -step
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
