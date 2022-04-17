package _5

import (
	"testing"
	"unicode"
)

/*
给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果NestedInteger 。
列表中的每个元素只可能是整数或整数嵌套列表
示例 1：
输入：s = "324",
输出：324
解释：你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
示例 2：
输入：s = "[123,[456,[789]]]",
输出：[123,[456,[789]]]
解释：返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
1. 一个 integer 包含值 123
2. 一个包含两个元素的嵌套列表：
    i.  一个 integer 包含值 456
    ii. 一个包含一个元素的嵌套列表
         a. 一个 integer 包含值 789
*/
func TestName(t *testing.T) {

}

type NestedInteger struct {
}

func (n *NestedInteger) Add(elem NestedInteger) {}
func (n *NestedInteger) SetInteger(value int)   {}
func deserialize(s string) *NestedInteger {
	index := 0
	var dfs func() *NestedInteger
	dfs = func() *NestedInteger {
		ni := &NestedInteger{}
		if s[index] == '[' {
			index++
			for s[index] != ']' {
				ni.Add(*dfs())
				if s[index] == ',' {
					index++
				}
			}
			index++
			return ni
		}

		negative := s[index] == '-'
		if negative {
			index++
		}
		num := 0
		for ; index < len(s) && unicode.IsDigit(rune(s[index])); index++ {
			num = num*10 + int(s[index]-'0')
		}
		if negative {
			num = -num
		}
		ni.SetInteger(num)
		return ni
	}
	return dfs()
}
