/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */
package q00055

// @lc code=start
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}

// @lc code=end
