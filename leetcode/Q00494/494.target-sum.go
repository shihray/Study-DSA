/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */
package q00494

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	return Dfs(nums, len(nums) - 1, target)
}

func Dfs(nums []int, index int, sum int) int {

	if index == -1 {
		if sum == 0 {
			return 1
		}
		return 0
	}

	count := 0
	for i := -1; i <= 1; i+=2 {
		count += Dfs(nums, index-1, sum+ (i* nums[index]))
	}
	return count
} 
// @lc code=end

