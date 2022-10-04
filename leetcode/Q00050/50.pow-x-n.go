/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
package q00050

// @lc code=start
func myPow(x float64, n int) float64 {
    minus := n < 0
	if minus {
		n = -n
	}
	resp := float64(1)
	for n != 0 {
		if n & 1 != 0 {
			resp *= x 
		}
		x *= x
		n = n >> 1
	}
	if minus {
		resp = 1/ resp
	}
	return resp
}
// @lc code=end

