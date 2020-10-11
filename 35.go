package main

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if target < v {
			if i == len(nums)-1 && len(nums) != 1 {
				return i
			}
			if i == 0 {
				return 0
			}
			return i
		}
	}
	return len(nums)
}
