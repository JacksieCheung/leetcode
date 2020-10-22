package main

// 机器人运动范围
func movingCount(m int, n int, k int) int {
	x := 0
	y := 0
	count := 1
	nums := make([][]int, m)
	for i := 0; i < m; i++ {
		nums[i] = make([]int, n)
		for j := 0; j < n; j++ {
			nums[i][j] = 0
		}
	}
	nums[0][0] = 1
	up(k, x, y, m, n, nums, &count)
	right(k, x, y, m, n, nums, &count)
	return count
}

func right(k int, x int, y int, m int, n int, nums [][]int, count *int) {
	x = x + 1
	if x >= m || nums[x][y] == 1 {
		return
	}
	sum := 0
	for i, j := x, y; i != 0 || j != 0; {
		sum = sum + i%10 + j%10
		i /= 10
		j /= 10
	}
	if sum > k {
		return
	}
	*count += 1
	nums[x][y] = 1
	fmt.Println(x, y)
	up(k, x, y, m, n, nums, count)
	right(k, x, y, m, n, nums, count)
}

func up(k int, x int, y int, m int, n int, nums [][]int, count *int) {
	y = y + 1
	if y >= n || nums[x][y] == 1 {
		return
	}
	sum := 0
	for i, j := x, y; i != 0 || j != 0; {
		sum = sum + i%10 + j%10
		i /= 10
		j /= 10
	}
	if sum > k {
		return
	}
	*count++
	nums[x][y] = 1
	fmt.Println(x, y)
	right(k, x, y, m, n, nums, count)
	up(k, x, y, m, n, nums, count)
}
