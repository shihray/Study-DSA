/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */
package q00278

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool;

func firstBadVersion(n int) int {

	start, end, firstBadVersion := 1, n, n

	for start < end {
		mid := (start + end) / 2
		result := isBadVersion(mid)

		if result {
			end = mid
			firstBadVersion = mid
		} else {
			start = mid + 1
		}
	}
	return firstBadVersion
}
// @lc code=end

