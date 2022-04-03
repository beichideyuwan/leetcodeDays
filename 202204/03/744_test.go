package _3

import (
	"sort"
	"testing"
)

/*
弱智题
给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母target，请你寻找在这一有序列表里比目标字母大的最小字母。
在比较时，字母是依序循环出现的。举个例子：
如果目标字母 target = 'z' 并且字符列表为letters = ['a', 'b']，则答案返回'a'
*/

func TestNextGreatestLetter(t *testing.T) {
	t.Log(nextGreatestLetter([]byte{'a', 'b', 'c'}, 'b'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	l := len(letters) - 1
	if target >= letters[l] {
		return letters[0]
	}
	search := sort.Search(l, func(i int) bool { return letters[i] > target })
	return letters[search]
}
