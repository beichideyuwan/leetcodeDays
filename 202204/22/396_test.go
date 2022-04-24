package _2

/*
给定一个长度为 n 的整数数组nums。
假设arrk是数组nums顺时针旋转 k 个位置后的数组，我们定义nums的 旋转函数F为：
F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]
返回F(0), F(1), ..., F(n-1)中的最大值。
生成的测试用例让答案符合32 位 整数。
*/

func maxRotateFunction(nums []int) int {
	numSum := 0
	for _, v := range nums {
		numSum += v
	}
	f := 0
	for i, num := range nums {
		f += i * num
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
