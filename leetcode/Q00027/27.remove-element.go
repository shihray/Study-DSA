/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */
package q00027

import "sort"

// @lc code=start
func removeElement(nums []int, val int) int {
    count := 0 
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 101
			count++
		}
	}
	sort.Ints(nums)
	return len(nums)-count
}
// @lc code=end

