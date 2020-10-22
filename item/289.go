package main

// 生命游戏
func gameOfLife(board [][]int) [][]int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				if i-1 >= 0 && j-1 >= 0 {
					if board[i-1][j-1] == 1 || board[i-1][j-1] == -1 {
						count++
					}
				}
				if i-1 >= 0 {
					if board[i-1][j] == 1 || board[i-1][j] == -1 {
						count++
					}
				}
				if j-1 >= 0 {
					if board[i][j-1] == 1 || board[i][j-1] == -1 {
						count++
					}
				}
				if i+1 < len(board) && j-1 >= 0 {
					if board[i+1][j-1] == 1 || board[i+1][j-1] == -1 {
						count++
					}
				}
				if i+1 < len(board) {
					if board[i+1][j] == 1 || board[i+1][j] == -1 {
						count++
					}
				}
				if j+1 < len(board[0]) {
					if board[i][j+1] == 1 || board[i][j+1] == -1 {
						count++
					}
				}
				if i+1 < len(board) && j+1 < len(board[0]) {
					if board[i+1][j+1] == 1 || board[i+1][j+1] == -1 {
						count++
					}
				}
				if i-1 >= 0 && j+1 < len(board[0]) {
					if board[i-1][j+1] == 1 || board[i-1][j+1] == -1 {
						count++
					}
				}
				if count == 3 {
					board[i][j] = 2
				}
				count = 0
			} else {
				if i-1 >= 0 && j-1 >= 0 {
					if board[i-1][j-1] == 1 || board[i-1][j-1] == -1 {
						count++
					}
				}
				if i-1 >= 0 {
					if board[i-1][j] == 1 || board[i-1][j] == -1 {
						count++
					}
				}
				if j-1 >= 0 {
					if board[i][j-1] == 1 || board[i][j-1] == -1 {
						count++
					}
				}
				if i+1 < len(board) && j-1 >= 0 {
					if board[i+1][j-1] == 1 || board[i+1][j-1] == -1 {
						count++
					}
				}
				if i+1 < len(board) {
					if board[i+1][j] == 1 || board[i+1][j] == -1 {
						count++
					}
				}
				if j+1 < len(board[0]) {
					if board[i][j+1] == 1 || board[i][j+1] == -1 {
						count++
					}
				}
				if i+1 < len(board) && j+1 < len(board[0]) {
					if board[i+1][j+1] == 1 || board[i+1][j+1] == -1 {
						count++
					}
				}
				if i-1 >= 0 && j+1 < len(board[0]) {
					if board[i-1][j+1] == 1 || board[i-1][j+1] == -1 {
						count++
					}
				}
				if count > 3 || count < 2 {
					board[i][j] = -1
				}
				count = 0
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
				continue
			}
			if board[i][j] == -1 {
				board[i][j] = 0
				continue
			}
		}
	}
	return board
}
