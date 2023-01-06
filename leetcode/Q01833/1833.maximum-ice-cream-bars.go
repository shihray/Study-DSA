/*
 * @lc app=leetcode id=1833 lang=golang
 *
 * [1833] Maximum Ice Cream Bars
 */
package q01833

import "sort"

// @lc code=start
func maxIceCream(costs []int, coins int) int {

	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	resp := 0
	for i := 0; i < len(costs); i++ {
		coins = coins - costs[i]
		if coins >= 0 {
			resp++
		} else {
			break
		}
	}
	return resp
}

// @lc code=end
