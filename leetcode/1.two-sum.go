/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	
	for i := 0; i < len(nums); i++ {
		output := []int{i}
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				output = append(output, j)
				return output
			}
		}
	}
	return nil
}
// @lc code=end

