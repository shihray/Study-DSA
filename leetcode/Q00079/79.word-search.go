/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */
package q00079

// @lc code=start
func exist(board [][]byte, word string) bool {
    
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if isValid(board, i, j, []byte(word), 0) {
					return true
				}
			}
		}
	}
	return false
}

func isValid(board [][]byte, row, col int, word []byte, index int) bool {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || board[row][col] == '.' {
		return false
	}

	if board[row][col] != word[index] {
		return false
	}

	if index == len(word) - 1 {
		return true
	}

	index += 1
	tmp := board[row][col]

	board[row][col] = '.'

	result := isValid(board, row + 1, col, word, index) ||
	isValid(board, row - 1, col, word, index) ||
	isValid(board, row, col + 1, word, index) ||
	isValid(board, row, col - 1, word, index) 

	board[row][col] = tmp
	return result
}
// @lc code=end

