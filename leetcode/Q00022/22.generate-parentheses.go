/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package q00022

// @lc code=start
func generateParenthesis(n int) []string {

	resp := []string{}
    
	var backtrack func( str string, open, close int)

	backtrack = func( str string, open, close int) {
		if len(str) == 2 * n {
			resp = append(resp, str)
		}
		if open < n {
			backtrack(str+"(", open + 1, close)
		}
		if close < open {
			backtrack(str+")", open, close + 1)
		}
	}

	backtrack("", 0, 0)
	return resp
}
// @lc code=end

