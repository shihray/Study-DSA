/*
 * @lc app=leetcode id=907 lang=golang
 *
 * [907] Sum of Subarray Minimums
 */
package q00907

import "math"

// @lc code=start
func sumSubarrayMins(arr []int) int {
	resp := 0
	for i := 0; i < len(arr); i++ {
		resp += combine(arr, i)
	}
	return resp
}

func combine(arr []int, k int) int {
	resp := 0

	var comb func(cur []int) 
	comb = func(cur []int) {
		if len(cur) == k {
			min := math.MaxInt64
			for i := 0; i < len(cur); i++ {
				if min > cur[i] {
					min = cur[i]
				}
			}
		}
		for i := 0; i < len(arr); i++ {
			list := []int{}
			copy(list, cur[:i])
			comb(list)
			cur = cur[:len(cur)-1]
		}
	}

	comb(arr) 
	return resp 
}

// @lc code=end

