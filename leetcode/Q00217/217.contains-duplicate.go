/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */
package q00217

// @lc code=start
func containsDuplicate(nums []int) bool {
    dupMap := map[int]int{}
	for _, val := range nums {
		if _, ok := dupMap[val]; ok {
			return true
		}
		dupMap[val]++
	}
	return false
}
// @lc code=end

