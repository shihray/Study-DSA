/*
 * @lc app=leetcode id=205 lang=golang
 *
 * [205] Isomorphic Strings
 */

// need to maintain same byte different place

// @lc code=start
func isIsomorphic(s string, t string) bool {
    sPattern, tPattern := map[uint8]int{}, map[uint8]int{}
	for index := range s {
		if sPattern[s[index]] != tPattern[t[index]] {
			return false
		} else {
			sPattern[s[index]] = index + 1 
			tPattern[t[index]] = index + 1
		}
	}

	return true
}
// @lc code=end

