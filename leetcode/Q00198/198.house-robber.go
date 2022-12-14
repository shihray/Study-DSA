/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */
package q00198

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	r := make([]int, n)
	r[0] = nums[0]
	r[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		r[i] = max(r[i-2]+nums[i], r[i-1])
	}
	return r[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

