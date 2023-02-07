/*
 * @lc app=leetcode id=1470 lang=golang
 *
 * [1470] Shuffle the Array
 */
package q01470

// @lc code=start
func shuffle(nums []int, n int) []int {

	resp := make([]int, len(nums))
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		resp[j], resp[j+1] = nums[i], nums[i+n]
	}
	return resp
}

// @lc code=end
