package _0220321

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

/*
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。
示例 1:
输入: [10,2]
输出: "102"
示例 2:
输入: [3,30,34,5,9]
输出: "3033459"
提示:
0 < nums.length <= 100
说明:
输出结果可能非常大，所以你需要返回一个字符串而不是整数
拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0
*/
func TestMinNumber(t *testing.T) {
	number := minNumber([]int{10, 23, 111, 321, 999})
	t.Log(number)
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x:=strconv.Itoa(nums[i])+strconv.Itoa(nums[j])
		y:=strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
		if x>y{
			return false
		}
		return true
	})
	ans :=""
	for _, num := range nums {
		ans+=fmt.Sprintf("%d",num)
	}
	return ans
}