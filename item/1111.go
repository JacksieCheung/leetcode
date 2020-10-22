package main

// 有效括号的嵌套深度
func maxDepthAfterSplit(seq string) []int {
	stack := make([]int, 0, len(seq))
	var deep = -1
	for _, v := range seq {
		if string(v) == "(" {
			deep++
			stack = append(stack, deep%2)
		}
		if string(v) == ")" {
			stack = append(stack, deep%2)
			deep--
		}
	}
	return stack
}
