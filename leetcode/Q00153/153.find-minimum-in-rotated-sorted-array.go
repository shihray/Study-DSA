/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */
package q00153

import "sort"

// @lc code=start
func findMinSort(nums []int) int {
    sort.Ints(nums)
	return nums[0]
}

func findMin(nums []int) int {
    start, end := 0, len(nums)-1

	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}
// @lc code=end

