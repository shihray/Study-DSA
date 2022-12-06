/*
 * @lc app=leetcode id=1704 lang=golang
 *
 * [1704] Determine if String Halves Are Alike
 */
package q01704

// @lc code=start
func halvesAreAlike(s string) bool {
    l := len(s) / 2 

	m := map[byte]bool{
		'a':false, 'e':false, 'i':false, 'o':false, 'u':false, 
		'A':false, 'E':false, 'I':false, 'O':false, 'U':false,
	}
	count := 0

	for i := 0; i < l; i++ {
		if _, exist := m[s[i]]; exist {
			count++
		}
	}
	for i := l; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			count--
		}
	}
	return count == 0
}
// @lc code=end

