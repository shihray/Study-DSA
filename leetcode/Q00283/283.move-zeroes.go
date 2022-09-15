/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */
package q00283

// @lc code=start
func moveZeroes(nums []int)  {
	writePoint := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			writePoint++
		} else if writePoint > 0 {
			tmp := nums[r]
			nums[r] = 0
			nums[r - writePoint] = tmp
		}
	}
}
// @lc code=end

