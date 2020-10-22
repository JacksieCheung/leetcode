package main

// 给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。
// 如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
func lengthOfLastWord(s string) int {
	l := len(s)
	count := -1
	l--
	for i := l; i >= 0; i-- {
		if i == l && s[i] == ' ' {
			l--
			continue
		}
		if s[i] == ' ' {
			return l - i
		}
		count++
		if count == l {
			return l + 1
		}
	}
	return 0
}
