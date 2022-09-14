/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package leetcode

// @lc code=start
func lengthOfLongestSubstring(s string) int {
    var resp int
	m := map[byte]int{}
	start := 0
	for end := 0; end < len(s); end++ {
		val, ok := m[s[end]]
		if ok {
			resp = max(resp, end-start)

			for i := start; i <= val; i++ {
				delete(m, s[i])
			}
			start = val + 1
		}
		m[s[end]] = end
	}

	resp = max(resp, len(s)-start)

	return resp
}


func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// @lc code=end

