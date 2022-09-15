/*
 * @lc app=leetcode id=1051 lang=golang
 *
 * [1051] Height Checker
 */
package q01051

import "sort"

// @lc code=start
func heightChecker(heights []int) int {
	sortH := make([]int, len(heights))
    copy(sortH, heights)
	sort.Ints(sortH)

	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sortH[i] {
			count++
		}
	}
	return count
}
// @lc code=end

