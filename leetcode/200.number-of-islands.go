/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
    
	result := 0
	for i, _ := range grid {
		for j, v2 := range grid[i] {
			if v2 == '1' {
				visit(i, j, &grid)
				result++
			}
		}
	}
	return result
}

func visit(row, col int, grid *[][]byte) {
	m, n := len(*grid), len((*grid)[0])

	if (row < 0) || (row >= m) || (col < 0) || (col >= n) || ((*grid)[row][col] == '0') {
		return
	}

	(*grid)[row][col] = '0'

	visit(row+1, col, grid)
	visit(row-1, col, grid)
	visit(row, col+1, grid)
	visit(row, col-1, grid)

}
// @lc code=end

