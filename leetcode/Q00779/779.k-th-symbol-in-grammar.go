/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */
package q00779

// @lc code=start
func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}

	if k%2 == 0 {
		return (1 - kthGrammar(n-1, (k+1)>>1))
	} else {
		return kthGrammar(n-1, (k+1)>>1)
	}
}

// @lc code=end
