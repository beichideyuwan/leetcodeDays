package _7

/*
字典树经典题型
单词数组words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：
words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给定一个单词数组words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。
*/

type Trie struct {
	child [26]*Trie
	depth int
}

func Constructor() *Trie {
	return &Trie{
		child: [26]*Trie{},
		depth: 1,
	}
}

func (t *Trie) Insert(word string) {
	curr := t
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		w := word[i] - 'a'
		if curr.child[w] == nil {
			curr.child[w] = Constructor()
		}
		curr.child[w].depth = curr.depth + 1
		curr = curr.child[w]
	}
}

func minimumLengthEncoding(words []string) int {
	t := Constructor()
	for _, word := range words {
		t.Insert(word)
	}
	var res int
	dfs := func(node *Trie) {}
	dfs = func(node *Trie) {
		var isChild bool
		for i := 0; i < 26; i++ {
			if node.child[i] != nil {
				isChild = true
				dfs(node.child[i])
			}
		}
		if !isChild {
			res += node.depth
		}
	}
	dfs(t)
	return res
}
