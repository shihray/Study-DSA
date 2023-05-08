/*
 * @lc app=leetcode id=1572 lang=golang
 *
 * [1572] Matrix Diagonal Sum
 */
package q01572

// @lc code=start
func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		res += mat[i][i] + mat[i][len(mat)-i-1]
	}
	if len(mat)%2 == 1 {
		res -= mat[len(mat)/2][len(mat)/2]
	}
	return res
}

// @lc code=end
