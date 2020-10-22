package main

// 括号生成
func generateParenthesis(n int) []string {
	result := make([]string, 0)
	//stack := make([]byte,n)
	tmp := make([]byte, 0)
	dfs(n, &result, 1, 0, tmp, '(')
	return result
}

func dfs(n int, result *[]string, countLeft int, countRight int, tmp []byte, signal byte) {
	tmp = append(tmp, signal)
	//fmt.Println(tmp)
	//fmt.Println(countLeft)
	if countLeft+1 <= n {
		dfs(n, result, countLeft+1, countRight, tmp, '(')
	}
	if countRight+1 <= n && countRight < countLeft {
		dfs(n, result, countLeft, countRight+1, tmp, ')')
	}
	if countRight == n && countLeft == n {
		*result = append(*result, string(tmp))
	}
	return
}
