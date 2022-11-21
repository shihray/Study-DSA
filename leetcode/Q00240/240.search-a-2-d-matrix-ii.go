/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */
package q00240

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1

	for i < n && j >= 0 {
		val := matrix[i][j]

		if val == target {
			return true
		} else if val > target {
			j--
		} else {
			i++
		}
	}
	return false
}

// @lc code=end

