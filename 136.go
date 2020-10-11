package main

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	var index, count int
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 0; i < len(nums); i++ {
		count = 0
		for n := 0; n < len(nums); n++ {
			if i == n {
				continue
			}
			if nums[i] == nums[n] {
				break
			}
			count++
			if count == len(nums)-1 {
				index = nums[i]
			}
		}
	}
	return index
}
