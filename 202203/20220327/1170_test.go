package _0220327

import "testing"

/*
定义一个函数f(s)，统计s 中（按字典序比较）最小字母的出现频次 ，其中 s是一个非空字符串。
例如，若s = "dcce"，那么f(s) = 2，因为字典序最小字母是"c"，它出现了2 次。
现在，给你两个字符串数组待查表queries和词汇表words 。对于每次查询queries[i] ，
需统计 words 中满足f(queries[i])< f(W)的 词的数目 ，W 表示词汇表words中的每个词。
请你返回一个整数数组answer作为答案，其中每个answer[i]是第 i 次查询的结果。

*/
func TestNumSmallerByFrequency(t *testing.T) {
	frequency := numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"})
	t.Log(frequency)
}
func numSmallerByFrequency(queries []string, words []string) []int {
	wordsFnNum := [11]int{}
	for _, word := range words {
		wordsFnNum[_numSmallerByFrequency(word)]++
	}
	var result = make([]int, len(queries), len(queries))
	for index, query := range queries {
		n := _numSmallerByFrequency(query)
		for cnt, val := range wordsFnNum {
			if n < cnt {
				result[index] += val
			}
		}
	}

	return result
}

//效率低的方法
func lowNumSmallerByFrequency(queries []string, words []string) []int {
	q, w := make([]int, len(queries)), make([]int, len(words))
	ans := make([]int, len(queries))
	for i, query := range queries {
		q[i] = _numSmallerByFrequency(query)
	}
	for i, word := range words {
		w[i] = _numSmallerByFrequency(word)
	}
	for i, _ := range q {
		for j, _ := range w {
			if q[i] < w[j] {
				ans[i]++
			}
		}
	}
	return ans
}

func _numSmallerByFrequency(str string) int {
	var dict = [26]int{}
	for _, r := range str {
		dict[r-'a'] += 1
	}

	for _, v := range dict {
		if v > 0 {
			return v
		}
	}
	return 0
}
