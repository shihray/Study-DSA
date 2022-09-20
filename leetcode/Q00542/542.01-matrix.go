/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */
package q00542

import "math"

// @lc code=start

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func updateMatrix(mat [][]int) [][]int {
	var queue [][2]int

    for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				mat[i][j] =  math.MaxInt64
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			x, y := cur[0] + dir[0] , cur[1] + dir[1]
			if x < 0 || x >= len(mat) || y < 0 || y >= len(mat[0]) || mat[x][y] <= mat[cur[0]][cur[1]] {
				continue
			  }
			queue = append(queue, [2]int{x, y})
			mat[x][y] = mat[cur[0]][cur[1]] + 1
		}
	}

	return mat
}

// @lc code=end

