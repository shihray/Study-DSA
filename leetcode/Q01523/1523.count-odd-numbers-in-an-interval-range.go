/*
 * @lc app=leetcode id=1523 lang=golang
 *
 * [1523] Count Odd Numbers in an Interval Range
 */
package q01523

// @lc code=start
func countOdds(low int, high int) int {
	return (high + 1)/2 - low / 2
}

// @lc code=end
