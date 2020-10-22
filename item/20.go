package main

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
func isValid(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}
	var a = []byte(s)
	var stack []byte
	for i := range a {
		if s[i] == '(' || a[i] == '[' || a[i] == '{' {
			stack = append(stack, a[i])
		} else {
			if len(stack)-1 < 0 {
				return false
			}
			if stack[len(stack)-1] == '(' && a[i] == ')' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == '[' && a[i] == ']' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == '{' && a[i] == '}' {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
