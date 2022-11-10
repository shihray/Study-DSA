/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */
package q00509

// @lc code=start
func fib(n int) int {
    if n < 2 {
		return n
	}
	resp := fib(n-1) + fib(n - 2)
	return resp
}
// @lc code=end

