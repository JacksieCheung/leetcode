package main

// 实现 int sqrt(int x) 函数。
// 计算并返回 x 的平方根，其中 x 是非负整数。
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		if i*i == x {
			return i
		}
		if i*i >= x {
			return i - 1
		}
	}
	return 0
}
