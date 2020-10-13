package main

// 旋转矩阵
func rotate(matrix [][]int) {
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		tmp[i] = make([]int, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			tmp[i][j] = matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = tmp[len(matrix)-1-j][i]
		}
	}
}
