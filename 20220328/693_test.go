package _0220328

import "testing"

/*
给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。
*/
func TestName(t *testing.T) {
	bits := hasAlternatingBits(5)
	t.Log(bits)
}

func hasAlternatingBits(n int) bool {
	var res = 2
	for ; n != 0; n /= 2 {
		cur := n % 2
		if cur == res {
			return false
		}
		res = cur
	}
	return true
}
