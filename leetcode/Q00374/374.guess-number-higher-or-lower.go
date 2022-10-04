/*
 * @lc app=leetcode id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */
package q00374
// @lc code=start
/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {

	start, end := 1, n

	for start <= end {
		mid := (start + end) / 2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1: 
			start = mid + 1
		default:
			return mid
		}
	}
	return 0
}
// @lc code=end

