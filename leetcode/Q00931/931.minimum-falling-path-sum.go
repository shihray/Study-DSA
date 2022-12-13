/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */
package q00931

import (
	"fmt"
	"math"
)

// @lc code=start
func minFallingPathSumTmp(matrix [][]int) int {
	tmp := map[string]int{}
	resp := math.MaxInt32

	for i := 0; i < len(matrix); i++ {
		resp = min(resp, bfs(matrix, 0, i, tmp))
	}
    return resp
}

func bfs(arr [][]int, row, col int, val map[string]int) int {
	if row >= len(arr) || row < 0 || col >= len(arr[0]) || col < 0 {
		return math.MaxInt64
	}

	if row == len(arr[0]) - 1 {
		return arr[row][col]
	}

	key := fmt.Sprintf("%d,%d", row, col)
	if v, exist := val[key]; exist {
		return v
	}

	aLeft := bfs(arr, row+1, col-1, val)
	a := bfs(arr, row+1, col, val)
	aRight := bfs(arr, row+1, col+1, val)

	resp := min(min(a, aLeft), aRight) + arr[row][col]
	val[key] = resp

	return resp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

