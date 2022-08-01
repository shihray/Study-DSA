/*
 * @lc app=leetcode id=1672 lang=golang
 *
 * [1672] Richest Customer Wealth
 */
package leetcode

// @lc code=start
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, v := range accounts {
		cur := 0
		for _, v2 := range v {
			cur += v2
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
// @lc code=end

