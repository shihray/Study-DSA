/*
 * @lc app=leetcode id=2348 lang=golang
 *
 * [2348] Number of Zero-Filled Subarrays
 */
package q02348

// @lc code=start
func zeroFilledSubarray(nums []int) int64 {
	
	var resp int64

	left := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			resp += int64(i - left +1)
		} else {
			left = i + 1
		}
	}

	return resp
}

// @lc code=end
