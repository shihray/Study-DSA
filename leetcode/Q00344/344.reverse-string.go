/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */
package q00344

// @lc code=start
func reverseString(s []byte)  {
	helper(s, 0)
}

func helper(s []byte, index int) {
	size := len(s)-1
	if index > size / 2 {
		return
	}
	s[index], s[size - index] = s[size - index], s[index]
	helper(s, index+1)
}

// @lc code=end

