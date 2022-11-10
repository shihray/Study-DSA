/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */
package q00779

import "strconv"

// @lc code=start
func kthGrammar(n int, k int) int {
    str := grammar("0", n)
	resp, _ := strconv.Atoi(string(str[k]))
	return resp
}

func grammar(row string, n int) string {
	if n == 0 {
		return row
	}
	cur := ""
	for i := len(row)-1; i >= 0; i-- {
		cur += string(row[i])
	}
	row += cur
	return grammar(row, n-1)
}

// @lc code=end

