package main

//　岛屿数量
func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				//递归函数
				count++
				search(grid, i, j)
			}
		}
	}
	return count
}

func search(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	search(grid, i, j-1)
	search(grid, i-1, j)
	search(grid, i+1, j)
	search(grid, i, j+1)
}
