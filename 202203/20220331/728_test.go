package _0220331

import (
	"strconv"
	"testing"
)

/*
728. 自除数
自除数是指可以被它包含的每一位数整除的数。
例如，128 是一个 自除数 ，因为128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。
自除数 不允许包含 0 。
给定两个整数left和right ，返回一个列表，列表的元素是范围[left, right]内所有的 自除数 。
*/

func TestSelfDividingNumbers(t *testing.T) {
	t.Log(selfDividingNumbers(412, 412))
}

func selfDividingNumbers(left int, right int) (ans []int) {
	for ; left <= right; left++ {
		if ok := isSelfDividing(left); ok {
			ans = append(ans, left)
		}
	}
	return
}

func isSelfDividing(i int) bool {
	sub := i

	l := len(strconv.Itoa(i)) - 1
	for ; l > 0; l-- {
		ten := pingfang(10, l)
		j := sub / ten
		if j == 0 {
			return false
		}
		if i%j != 0 {
			return false
		} else {
			sub -= ten * j
		}
	}
	if sub == 0 {
		return false
	}
	if i%sub != 0 {
		return false
	}
	return true
}

func pingfang(i, j int) int {
	ans := 1
	for ; j > 0; j-- {
		ans *= i
	}
	return ans
}
