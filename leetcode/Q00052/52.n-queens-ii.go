/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 */
package q00052

// @lc code=start
func totalNQueens(n int) int {
	var mat [][]bool
    for i :=0; i < n; i++ {
        mat = append(mat, make([]bool, n))
    }
    
    var count = 0
    helper(mat, 0, &count)
    
    return count
}

func helper(mat [][]bool, row int, count *int) {
	n := len(mat)

	if row == n {
		*count += 1
		return
	}

	for c := 0; c < n; c++ {
		
		if isValid(mat, row, c) {
			mat[row][c] = true

			helper(mat, row + 1, count)

			mat[row][c] = false
		}
	}
}

func isValid(mat [][]bool, row int, col int) bool {
	n := len(mat)

	for i := row -1; i >= 0; i-- {
		if mat[i][col] {
			return false
		}
	}

	for i, j := row -1, col-1; i >= 0 && j >= 0; i,j = i-1, j-1 {
		if mat[i][j] {
			return false
		}
	}

	for i, j := row -1, col+1; i >= 0 && j < n; i,j = i-1, j+1 {
		if mat[i][j] {
			return false
		}
	}
	return true
}

// @lc code=end

