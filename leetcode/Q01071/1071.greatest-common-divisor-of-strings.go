/*
 * @lc app=leetcode id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */
package q01071

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		str1, str2 = str2, str1
	}
	if str1[:len(str2)]!= str2 {
		return ""
	}
	return gcdOfStrings(str1[len(str2):], str2)
}

// @lc code=end
