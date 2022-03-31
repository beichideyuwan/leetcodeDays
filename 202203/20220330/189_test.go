package _0220330

import (
	"fmt"
	"testing"
)

/*
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
*/
func TestRotate(t *testing.T) {
	s := []int{3, 2, 1, 56, 421}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	j := len(nums)
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[j-i-1] = nums[j-i-1], nums[i]
	}
}
