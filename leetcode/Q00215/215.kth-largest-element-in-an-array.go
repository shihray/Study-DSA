/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
package q00215

import "sort"

// @lc code=start
func findKthLargest(nums []int, k int) int {
    sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}
// @lc code=end