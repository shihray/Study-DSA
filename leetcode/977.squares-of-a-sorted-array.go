/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */
package leetcode

import "sort"

// @lc code=start
func sortedSquares(nums []int) []int {
    for index := range nums {
		nums[index] = nums[index] * nums[index]
	}

	sort.Ints(nums)
	return nums
}
// @lc code=end

