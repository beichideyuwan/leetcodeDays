package _02204

/*
给你一个整数数组 nums，和一个整数 k 。
在一个操作中，您可以选择 0 <= i < nums.length 的任何索引 i 。
将 nums[i] 改为 nums[i] + x ，其中 x 是一个范围为 [-k, k] 的整数。
对于每个索引 i ，最多 只能 应用 一次 此操作。
nums的分数是nums中最大和最小元素的差值。
在对 nums 中的每个索引最多应用一次上述操作后，返回nums 的最低 分数 。
*/
func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		} else if num > maxNum {
			maxNum = num
		}
	}
	ans := maxNum - minNum - 2*k
	if ans > 0 {
		return ans
	}
	return 0
}
