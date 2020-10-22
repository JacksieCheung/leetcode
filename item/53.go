package main

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组
// （子数组最少包含一个元素），返回其最大和。
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var index = 1
	var count = 0
	var maxSum = nums[0]
	var nowSum = nums[0]
	for {
		if nowSum > maxSum {
			maxSum = nowSum
		}
		if count == len(nums)-1 {
			break
		}
		nowSum += nums[index]
		index++
		if index > len(nums)-1 {
			if nowSum > maxSum {
				maxSum = nowSum
			}
			count++
			nowSum = nums[count]
			index = count + 1
		}
	}
	return maxSum
}
