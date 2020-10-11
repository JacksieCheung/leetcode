package main

// 判断一个整数是否是回文数。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	len := 0
	var nums = make([]int, 0)
	for x != 0 {
		nums = append(nums, x%10)
		len++
		x /= 10
	}
	k := len - 1
	for i := 0; i < len/2; i++ {
		if nums[i] != nums[k] {
			return false
		}
		k--
	}
	return true
}
