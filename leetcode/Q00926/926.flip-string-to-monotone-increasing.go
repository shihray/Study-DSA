/*
 * @lc app=leetcode id=926 lang=golang
 *
 * [926] Flip String to Monotone Increasing
 */
package q00926

// @lc code=start
func minFlipsMonoIncr(s string) int {
	occ_of_1 := 0
	flip := 0

	for _, char := range s {

		if char == '1' {
			occ_of_1 += 1
		} else {
			flip = min(flip+1, occ_of_1)
		}
	}
	return flip
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
