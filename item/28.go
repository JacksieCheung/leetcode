package main

// 给定一个 haystack 字符串和一个 needle 字符串，
// 在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
// 如果不存在，则返回  -1。
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if len(needle)+i <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
