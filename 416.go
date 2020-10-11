package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97}))
}

func canPartition(nums []int) bool {
	if len(nums) == 0 || nums == nil {
		return true
	}
	var leftSum int
	for _, v := range nums {
		leftSum += v
	}
	if leftSum%2 != 0 {
		return false
	}
	var curSum int
	var flag bool
	for i, v := range nums {
		if v > leftSum/2 {
			return false
		}
		numstmp := append([]int{}, nums...)
		dfs(numstmp, i, curSum+v, leftSum-v, &flag)
	}
	return flag
}

func dfs(nums []int, i int, curSum int, leftSum int, flag *bool) {
	if *flag == true {
		return
	}

	//	fmt.Println(nums, i, curSum, leftSum)
	if i == len(nums)-1 {
		//fmt.Println("hello", nums[:i])
		nums = nil
	} else {
		nums = append(nums[:i], nums[i+1:]...)
	}

	//fmt.Println(nums)
	if curSum > leftSum || len(nums) == 0 || nums == nil {
		return
	}
	if curSum == leftSum {
		*flag = true
		return
	}
	for i, v := range nums {
		numstmp := append([]int{}, nums...)
		dfs(numstmp, i, curSum+v, leftSum-v, flag)
	}
}
