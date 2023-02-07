/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */
package q00567

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	need := [26]int{}
	have := [26]int{}

	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
	}

	for i := 0; i < len(s2); i++ {
		
		if i >= len(s1) {
			have[s2[i-len(s1)]-'a']--
		}

		have[s2[i]-'a']++

		if need == have {
			return true
		}
	}
	return false
}

// @lc code=end
