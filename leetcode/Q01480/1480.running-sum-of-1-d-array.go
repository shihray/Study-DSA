/*
 * @lc app=leetcode id=1480 lang=golang
 *
 * [1480] Running Sum of 1d Array
 */

package q01480

// @lc code=start
func runningSum(nums []int) []int {
    for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
// @lc code=end

// tc = O(n)
// sc = O(1)
