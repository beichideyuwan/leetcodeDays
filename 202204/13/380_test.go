package _3

import "math/rand"

/*
实现RandomizedSet 类：
RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
*/

type RandomizedSet struct {
	nums []int
	ins  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{nums: []int{}, ins: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.ins[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.ins[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.ins[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last]
	this.ins[this.nums[id]] = id
	this.nums = this.nums[:last]
	delete(this.ins, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
