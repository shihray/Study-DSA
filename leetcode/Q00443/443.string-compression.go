/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */
package q00443

import "strconv"

// @lc code=start
func compress(chars []byte) int {

	slow := 0
	for i := 0; i < len(chars); {
		cnt := 1
		for i+1 < len(chars) && chars[i] == chars[i+1] {
			cnt++
			i++
		}

		chars[slow] = chars[i]
		slow++
		if cnt > 1 {
			for _, c := range strconv.Itoa(cnt) {
				chars[slow] = byte(c)
				slow++
			}
		}
		i++
	}
	return slow
}

// @lc code=end
