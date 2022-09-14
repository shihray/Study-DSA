/*
 * @lc app=leetcode id=1342 lang=golang
 *
 * [1342] Number of Steps to Reduce a Number to Zero
 */
package leetcode

// @lc code=start
func numberOfSteps(num int) int {
    count := 0

	for {
		if num == 0 {
			break
		}
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		count++
	}
	return count
}
// @lc code=end

