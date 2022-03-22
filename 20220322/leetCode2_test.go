package _0220322

import (
	"fmt"
	"testing"
)
/*
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
*/
func TestContainsDuplicate(t *testing.T) {
	fmt.Println(containsDuplicate([]int{1,2,3,4}))
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
}
func containsDuplicate(nums []int) bool {
	m:= map[int]struct{}{}
	for _, num := range nums {
		if _,ok:=m[num];!ok{
			m[num]= struct{}{}
		}else {
			return true
		}
	}
	return false
}