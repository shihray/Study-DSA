/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */
package leetcode

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	a, b := []rune(ransomNote), []rune(magazine)

	countArr := make([]int, 26)
	for i := 0; i < len(b); i++ {
		countArr[b[i] - 'a']++
	}
	for i := 0; i < len(a); i++ {
		countArr[a[i] - 'a']--
	}
	for i := 0; i < len(countArr); i++ {
		if countArr[i] < 0 {
			return false
		}
	}
	return true
}
// @lc code=end

