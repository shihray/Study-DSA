/*
 * @lc app=leetcode id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */
package q01657

import "sort"

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
		return false
	}

	k1, v1 := countArr(word1)
	k2, v2 := countArr(word2)
	return k1 == k2 && v1 == v2
}

func countArr(str string) (k, v [26]int) {
	
	for i := 0; i < len(str); i++ {
		k[str[i] - 'a'] = 1
		v[str[i] - 'a']++
	}
	sort.Ints(v[:])
	return
}
// @lc code=end

