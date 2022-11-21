/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */
package q00037

// @lc code=start
func solveSudoku(board [][]byte)  {
    solve(board, 0, 0)
}

func solve(board [][]byte, row, col int) bool {
	solved, row, col := getNetPos(board, row, col)
	if solved {
		return solved
	}

	for i := byte('1'); i <= '9'; i++ {
		if isValid(board, i, row, col) {
			add(board, row, col, i)

			solved := solve(board, row, col + 1)
			if solved {
				return solved
			}
			remove(board, row, col)
		}
	}
	return false
}

func getNetPos(board [][]byte, row, col int) (bool, int, int) {
	for row < 9 {
		for col < 9 {
			if board[row][col] == '.' {
				return false, row, col
			}
			col++
		}
		row++
		col = 0
	}
	return true, row, col
}

func remove(board [][]byte, row, col int) {
	board[row][col] = '.'
}

func add(board [][]byte, row, col int, val byte) {
	board[row][col] = val
}

func isValid(board [][]byte, val byte, row, col int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == val {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == val {
			return false
		}
	}

	rowStart, rowEnd := (row/3)*3, (row/3+1)*3
	colStart, colEnd := (col/3)*3, (col/3+1)*3

	for i := rowStart; i < rowEnd; i++ {
		for j := colStart; j < colEnd; j++ {
			if board[i][j] == val {
				return false
			}
		}
	}
	return true
}

// @lc code=end

