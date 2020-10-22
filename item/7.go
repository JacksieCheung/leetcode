package main

// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
func reverse(x int) int {
	var y int
	for x != 0 {
		y = y*10 + x%10
		if y > 1<<31-1 || y < (-1)<<31 {
			return 0
		}
		x /= 10
	}
	return y
}
