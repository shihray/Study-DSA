/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */
package q01143

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
    l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1 + 1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, l2 + 1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i][j-1] >dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[l1][l2]
}
// @lc code=end

