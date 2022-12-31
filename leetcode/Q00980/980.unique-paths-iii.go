/*
 * @lc app=leetcode id=980 lang=golang
 *
 * [980] Unique Paths III
 */
package q00980

// @lc code=start
func uniquePathsIII(grid [][]int) int {

	startX, startY, count := 0, 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				count++
			} else if grid[i][j] == 1 {
				startX, startY = i, j
			}
		}
	}

	return dfs(startX, startY, count, grid)
}

func dfs(row, col int, count int, grid [][]int) int {

	resp := 0
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] == -1 {
		return 0
	} else if grid[row][col] == 0 || grid[row][col] == 1 {
		grid[row][col] = -1

		resp = dfs(row+1, col, count-1, grid) +
			dfs(row-1, col, count-1, grid) +
			dfs(row, col+1, count-1, grid) +
			dfs(row, col-1, count-1, grid)

		grid[row][col] = 0

	} else if grid[row][col] == 2 && count < 0 {
		return 1
	}
	return resp
}

// @lc code=end
