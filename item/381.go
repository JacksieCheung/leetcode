package main

type RandomizedCollection struct {
	Nums      []int
	Length    int
	IndexHash map[int]map[int]struct{} // 第二个map只用到了key，val没用
	// val->map->indexs
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{IndexHash: make(map[int]map[int]struct{})}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	// fmt.Println("insert begin:",this.IndexHash)
	indexMap, ok := this.IndexHash[val]
	this.Nums = append(this.Nums, val) // 不管有没有重复都插入
	this.Length++
	if !ok { // 没有找到，插入的第一个
		indexMap = make(map[int]struct{})
		// indexMap=map[int]struct{}{}
		this.IndexHash[val] = indexMap // 顶层的 map 和底层 map 建立联系
	}
	indexMap[this.Length-1] = struct{}{} // 记录当前元素 val 插入索引
	// fmt.Println("Insert finished:",this.IndexHash)
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	// fmt.Println(this.Nums)
	// 先查找是否存在
	indexMap, ok := this.IndexHash[val]
	// fmt.Println("delete begin:",this.IndexHash)
	if len(indexMap) == 0 || !ok {
		return false
	}

	var delIndex int // 建立要删除的索引，用 range 一次找到目标索引
	for key, _ := range indexMap {
		delIndex = key
		break
	}

	this.Nums[delIndex] = this.Nums[this.Length-1]
	delete(indexMap, delIndex)                    // 删掉要删除的 val 对应的索引
	tmpMap := this.IndexHash[this.Nums[delIndex]] // 找到尾巴 val 所对应的 map
	delete(tmpMap, this.Length-1)                 // 对尾巴 val 所对应的 map 进行更改，将原来在末尾的删掉
	if delIndex < this.Length-1 {
		tmpMap[delIndex] = struct{}{} // 将新的索引位置输入 map
	}
	this.Nums = this.Nums[:this.Length-1]
	this.Length--
	// fmt.Println("delete finished:",this.IndexHash)
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	// fmt.Println(this.Length)
	if this.Length == 0 {
		return -1
	}
	return this.Nums[rand.Intn(this.Length)]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
