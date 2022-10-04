/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */
package studydsa

// @lc code=start
func mySqrt(x int) int {
    resp := 0 

	for resp * resp < x{
		resp++
	}

	if resp * resp > x {
		resp--
	}
	return resp
}
// @lc code=end

