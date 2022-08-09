/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
package leetcode

// @lc code=start
func min(a int, b int) int {
    if a < b {
        return a
    }
	return b
}

func binom(n int, k int) int {
    if k == 0 {
        return 1
    }
    return (n - k + 1) * binom(n, k - 1) / k
}

func uniquePaths(m int, n int) int {
    return binom(m + n - 2, min(m - 1, n - 1))
}
// @lc code=end

