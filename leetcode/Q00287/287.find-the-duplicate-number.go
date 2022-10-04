/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */
package q00287

// @lc code=start
func findDuplicate(nums []int) int {
    slow := nums[0]
    fast := nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    fast = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    return fast
}
// @lc code=end

