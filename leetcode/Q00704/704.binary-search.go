/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */
package studydsa

// @lc code=start
func search(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return i
        }
    }
    return -1
}
// @lc code=end

