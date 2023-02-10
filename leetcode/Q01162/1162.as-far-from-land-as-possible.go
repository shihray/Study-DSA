/*
 * @lc app=leetcode id=1162 lang=golang
 *
 * [1162] As Far from Land as Possible
 */
package q01162

// @lc code=start
func maxDistance(grid [][]int) int {
	direction, setps, ok := [][]int{}, [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}, false

	visited := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				direction = append(direction, []int{i, j})
			} else {
				ok = true
			}
		}
	}

	resp := -1
	if !ok {
		return resp
	}

	resp = 0
	for len(direction) > 0 {
		resp++
		for l := len(direction); l > 0; l-- {
			n := direction[0]
			direction = direction[1:]

			for _, diff := range setps {
				if i, j := n[0]+diff[0], n[1]+diff[1]; 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0]) {
					if grid[i][j] != 0 {
						continue
					}
					if visited[i][j] == 0 {
						visited[i][j] = resp
						direction = append(direction, []int{i, j})
					}
				}
			}

		}
	}
	return resp - 1
}

// @lc code=end
