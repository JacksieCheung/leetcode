package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97}))

}

var flag bool

func canPartition(nums []int) bool {
	flag = false
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return flag
	} else {

		sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		for i := 0; i < len(nums); i++ {
			if nums[i] > sum/2 {
				return false
			}
			if nums[i] < sum/2 {
				break
			}
		}
		do(sum/2, nums, 0, 0)
		return flag
	}
}
func do(sum int, nums []int, current, start int) {
	if current > sum {
		return
	}
	if current == sum {
		flag = true
		return
	}
	if !flag {
		for i := start; i < len(nums); i++ {
			if current+nums[i] <= sum {
				current += nums[i]
				do(sum, nums, current, i+1)
				current -= nums[i]
			}
		}
	}
}
