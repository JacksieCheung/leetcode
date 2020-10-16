package main

// 有序数组的平方
// 用归并排序 还能用双指针
func sortedSquares(A []int) []int {
	if len(A) == 0 {
		return A
	}
	var index int
	for i := 0; i < len(A); i++ {
		if A[i] >= 0 {
			index = i - 1
			break
		}
	}
	var res []int
	var pre, nex int
	for pre, nex = index, index+1; nex < len(A) || pre >= 0; {
		if nex < len(A) && pre >= 0 {
			if A[pre]*A[pre] > A[nex]*A[nex] {
				res = append(res, A[nex]*A[nex])
				nex++
			} else {
				res = append(res, A[pre]*A[pre])
				pre--
			}
		} else {
			if nex < len(A) {
				res = append(res, A[nex]*A[nex])
				nex++
			}
			if pre >= 0 {
				res = append(res, A[pre]*A[pre])
				pre--
			}
		}
	}
	return res
}
