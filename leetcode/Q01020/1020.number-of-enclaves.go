/*
 * @lc app=leetcode id=1020 lang=golang
 *
 * [1020] Number of Enclaves
 */
package q01020

// @lc code=start
func numEnclaves(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
            dfs(grid,i,0);
        }
        if grid[i][len(grid[i])-1] == 1 {
            dfs(grid,i,len(grid[i])-1)
        }
	}

	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 1 {
            dfs(grid,0,i);
        }
        if grid[len(grid)-1][i] == 1 {
            dfs(grid,len(grid)-1,i);
        }
	}

	resp := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				resp += dfs(grid, i, j)
			}
		}
	}
	return resp
}

func dfs(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		return 0
	}

	if grid[x][y] == 0 {
		return 0
	}

	grid[x][y] = 0
	top := dfs(grid, x-1, y)
	down := dfs(grid, x+1, y)
	left := dfs(grid, x, y-1)
	right := dfs(grid, x, y+1)

	return 1 + top + down + left + right
}

// @lc code=end
