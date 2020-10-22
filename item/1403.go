package main

// 非递增顺序最小子序列
func minSubsequence(nums []int) []int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	res := 0
	last := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res += nums[i]
		if res > sum-res {
			last = i
			break
		}
	}

	var result []int
	for i := len(nums) - 1; i >= last; i-- {
		result = append(result, nums[i])
	}
	return result
}
