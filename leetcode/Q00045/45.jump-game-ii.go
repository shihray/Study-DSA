/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */
package q00045

// @lc code=start
func jump(nums []int) int {

	resp := 0
	curEnd, curFar := 0, 0

	for i := 0; i < len(nums)-1; i++ {
		curFar = max(curFar, i + nums[i])

		if i == curEnd {
			resp++
			curEnd = curFar
		}
	}

	return resp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
