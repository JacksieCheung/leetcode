package main

// 递归实现，超时
func dfs(n int, count *int) {
	if n == 0 {
		(*count)++
		return
	}
	if n >= 1 {
		dfs(n-1, count)
	}
	if n >= 2 {
		dfs(n-2, count)
	}
}

func climbStairs(n int) int {
	count := 0
	dfs(n, &count)
	return count
}
