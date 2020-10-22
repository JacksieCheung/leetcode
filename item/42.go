package main

// 接雨水
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	sum := 0
	max := 0
	for i := 0; i < len(height); i++ {
		if height[i] > height[max] {
			max = i
		}
	}
	left := height[0]
	for i := 0; i < max; i++ {
		if left > height[i] {
			sum += left - height[i]
		} else {
			left = height[i]
		}
	}
	right := height[len(height)-1]
	for i := len(height) - 1; i > max; i-- {
		if right > height[i] {
			sum += right - height[i]
		} else {
			right = height[i]
		}
	}
	return sum
}
