/*
 * @lc app=leetcode id=944 lang=golang
 *
 * [944] Delete Columns to Make Sorted
 */
package q00944

// @lc code=start
func minDeletionSize(strs []string) int {

	resp := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				resp++
				break
			}
		}
	}
	return resp
}

// @lc code=end
