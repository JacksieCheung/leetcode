package main

// 要构造一个数据结构使得插入删除随机获取都是 O(1)
type RandomizedSet struct {
	Nums      []int
	Length    int         // 存储数组长度
	IndexHash map[int]int // 保存数组索引的结构，key是数组值，valude是数组索引
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{IndexHash: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.IndexHash[val]
	if !ok {
		this.Nums = append(this.Nums, val)
		this.Length++
		this.IndexHash[val] = this.Length - 1
		return true
	}
	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	elemIndex, ok := this.IndexHash[val]
	if ok {
		this.Nums[elemIndex] = this.Nums[this.Length-1] // 这里可以直接覆盖掉，并不需要交换
		this.IndexHash[this.Nums[elemIndex]] = elemIndex
		// 先变换再删除，防止最后一个元素造成 index out of range
		delete(this.IndexHash, val)
		this.Nums = this.Nums[:this.Length-1]
		this.Length--
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.Length == 0 {
		return -1
	}
	index := rand.Intn(this.Length)
	return this.Nums[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
