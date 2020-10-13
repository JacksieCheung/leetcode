package main

// 反转字符串里单词
func reverseWords(s string) string {
	//双指针遍历字符串，当第二个指到空格，把这一段切出来。然后前面的指针等于后面的继续遍历
	if len(s) == 1 && s != " " {
		return s
	}
	str := make([]string, 0)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(i, j, len(s)-1)
			if s[i] == ' ' {
				break
			}
			if s[j] == ' ' {
				str = append(str, s[i:j])
				i = j
				i--
				break
			}
			if j == len(s)-1 {
				str = append(str, s[i:])
				i = j
				break
			}
		}
	}

	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += str[i]
		if i != 0 {
			result += " "
		}
	}
	return result
}
