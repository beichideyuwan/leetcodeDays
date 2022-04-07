package _5

import (
	"math/bits"
	"testing"
)

/*
给你两个整数left和right ，在闭区间 [left, right]范围内，统计并返回 计算置位位数为质数 的整数个数。
计算置位位数 就是二进制表示中 1 的个数。
例如， 21的二进制表示10101有 3 个计算置位。
*/
func TestCountPrimeSetBits(t *testing.T) {
	t.Log(countPrimeSetBits(2, 19))
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeSetBits(left, right int) (ans int) {
	for x := left; x <= right; x++ {
		if isPrime(bits.OnesCount(uint(x))) {
			ans++
		}
	}
	return
}
