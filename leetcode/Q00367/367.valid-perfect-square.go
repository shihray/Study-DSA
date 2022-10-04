/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

package q00367

// @lc code=start
func isPerfectSquare(num int) bool {
	start, end := 1, num

	for start <= end {
		mid := (start + end) / 2
		sqrt := mid * mid
		if sqrt == num {
			return true
		} else if sqrt > num {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
// @lc code=end

