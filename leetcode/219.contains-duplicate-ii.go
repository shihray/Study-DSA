/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

package leetcode

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int][]int{}
	for i := len(nums) - 1; i >= 0; i-- {

		for _, val := range m[nums[i]] {
			if val-i <= k {
				return true
			}
		}
		m[nums[i]] = append(m[nums[i]], i)
	}
	return false
}
// @lc code=end

