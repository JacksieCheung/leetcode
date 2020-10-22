package main

func sortedSquares(A []int) []int {
	nex := len(A) - 1
	pre := 0
	res := make([]int, nex+1)
	for pos := len(A) - 1; pos >= 0; pos-- {
		if v, w := A[pre]*A[pre], A[nex]*A[nex]; v > w {
			res[pos] = v
			pre++
		} else {
			res[pos] = w
			nex--
		}
	}
	return res
}
