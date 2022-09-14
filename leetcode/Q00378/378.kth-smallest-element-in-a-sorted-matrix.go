/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */
package leetcode

import "sort"

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
    
	var resp []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			resp = append(resp, matrix[i][j])
		}
	}

	sort.Ints(resp)
	return resp[k-1]
}
// @lc code=end

