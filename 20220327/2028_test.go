package _0220327

import "testing"

func TestMissingRolls(t *testing.T) {
	t.Log(missingRolls([]int{1, 5, 6}, 3, 4))
}

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	p := mean * (n + m)
	ans := make([]int, n)
	for _, roll := range rolls {
		p -= roll
	}
	if p > 6*n || p < n {
		return []int{}
	}
	s := p / n
	b := 6 - s
	a := p % n
	for i := 0; i < n; i++ {
		if a > b {
			ans[i] = b
			a -= b
		} else if a != 0 {
			ans[i] = a
			a -= a
		}
		ans[i] += s
	}
	return ans
}
