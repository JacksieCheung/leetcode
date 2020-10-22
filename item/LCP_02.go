package main

// 分式化简
func fraction(cont []int) []int {
	if len(cont) == 1 {
		return []int{cont[0], 1}
	}
	v := fraction(cont[1:])
	return []int{cont[0]*v[0] + v[1], v[0]}
}
