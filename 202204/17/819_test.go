package _7

import (
	"strings"
	"unicode"
)

/*
给定一个段落 (paragraph) 和一个禁用单词列表 (banned)。返回出现次数最多，同时不在禁用列表中的单词。
题目保证至少有一个词不在禁用列表中，而且答案唯一。
禁用列表中的单词用小写字母表示，不含标点符号。段落中的单词不区分大小写。答案都是小写字母。
*/
func mostCommonWord(paragraph string, banned []string) string {
	// 切分单词，转小写，遍历max及banned
	// 非字母切分段落
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	mapWords := make(map[string]int)
	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		mapWords[word]++
	}
	max := 0
	res := ""
	for word, count := range mapWords {
		// 数量达且不在banned中的
		if count > max && !isBanned(word, banned) {
			max = count
			res = word
		}
	}
	return res
}

func isBanned(word string, banned []string) bool {
	for _, bannedString := range banned {
		if word == bannedString {
			return true
		}
	}
	return false
}
