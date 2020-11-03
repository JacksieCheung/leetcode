package main

// 有效山脉
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	maxIndex := 0
	count := 0
	for i, v := range A {
		if A[maxIndex] <= v {
			if A[maxIndex] == v {
				count++
				continue
			}
			maxIndex = i
			count = 1
		}
	}

	if count != 1 {
		return false
	}

	if maxIndex == 0 || maxIndex == len(A)-1 {
		return false
	}

	i := maxIndex - 1
	j := maxIndex + 1
	fmt.Println(i, j)
	for {
		if i >= 0 {
			if A[i] >= A[i+1] {
				return false
			}
			i--
		}
		if j <= len(A)-1 {
			if A[j-1] <= A[j] {
				return false
			}
			fmt.Println("hello")
			j++
		}
		if i == -1 && j == len(A) {
			break
		}
	}
	return true
}
