package main

// 给定一个整数数组 nums 和一个目标值 target，
// 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
func twoSum(nums []int, target int) []int {

	var numMap map[int]int
	numMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		var difference = target - nums[i]
		if v, ok := numMap[difference]; ok {
			return []int{v, i}
		} else {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}
