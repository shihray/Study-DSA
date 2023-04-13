/*
 * @lc app=leetcode id=1254 lang=golang
 *
 * [1254] Number of Closed Islands
 */
package q01254

// @lc code=start
func closedIsland(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && (i == 0 || i == m-1 || j == 0 || j == n-1) {
				dfs(grid, i, j)
			}
		}
	}

	var islands int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dfs(grid, i, j)
				islands++
			}
		}
	}
	return islands
}

func dfs(grid [][]int, x, y int) {
	if x < 0 || y < 0 || x == len(grid) || y == len(grid[0]) || grid[x][y] == 1 {
		return
	}

	grid[x][y] = 1
	dfs(grid, x+1, y)
	dfs(grid, x-1, y)
	dfs(grid, x, y+1)
	dfs(grid, x, y-1)
}

// @lc code=end
