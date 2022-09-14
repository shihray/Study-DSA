/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */
package leetcode

// @lc code=start
func sortArrayByParity(nums []int) []int {
    begin, end := 0, len(nums)-1
    
    for begin <= end {
        if nums[begin] % 2 == 0 {
            begin++
        } else {
            nums[begin], nums[end] = nums[end], nums[begin]
            end--
        }
    }
    return nums
}
// @lc code=end

