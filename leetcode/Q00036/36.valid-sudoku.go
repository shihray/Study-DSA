/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */
package q00036

import "strconv"

// @lc code=start
func isValidSudoku(board [][]byte) bool {
    rowArr, colArr, gridArr := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val, err := strconv.Atoi(string(board[i][j]))
            if err != nil {
                continue
            }

			val--
            gridIdx := j/3 + (i/3) * 3
			if rowArr[i][val] || colArr[j][val] || gridArr[gridIdx][val] {
				return false
			}

			rowArr[i][val] = true
			colArr[j][val] = true
			gridArr[gridIdx][val] = true
		}
	}
	return true
}
// @lc code=end

