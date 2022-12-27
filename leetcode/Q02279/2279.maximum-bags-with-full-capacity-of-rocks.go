/*
 * @lc app=leetcode id=2279 lang=golang
 *
 * [2279] Maximum Bags With Full Capacity of Rocks
 */
package q02279

import "sort"

// @lc code=start
func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	size := make([]int, len(capacity))

	for i := 0; i < len(capacity); i++ {
		size[i] = capacity[i] - rocks[i]
	}

	sort.Ints(size)

	resp := 0
	for _, v := range size {
		if additionalRocks-v >= 0 {
			additionalRocks = additionalRocks - v
			resp++
		}
	}
	return resp
}

// @lc code=end
