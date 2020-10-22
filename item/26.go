package main

// 给定一个排序数组，你需要在 原地 删除重复出现的元素，
// 使得每个元素只出现一次，返回移除后数组的新长度。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i, k int
	i = 0
	for k = 1; k < len(nums); k++ {
		if nums[i] != nums[k] {
			i++
			nums[i] = nums[k]
		}
	}
	return i + 1
}
