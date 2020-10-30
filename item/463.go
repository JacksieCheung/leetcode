package main

// 岛屿周长 用遍历
func islandPerimeter(grid [][]int) int {
	length := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 1 {
				if x-1 < 0 || grid[x-1][y] == 0 {
					length++
				}
				if x+1 >= len(grid) || grid[x+1][y] == 0 {
					length++
				}
				if y-1 < 0 || grid[x][y-1] == 0 {
					length++
				}
				if y+1 >= len(grid[x]) || grid[x][y+1] == 0 {
					length++
				}
			}
		}
	}
	return length
}
