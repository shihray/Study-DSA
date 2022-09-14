/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

package leetcode

// @lc code=start
func groupAnagrams(strs []string) [][]string {
    
	mp := map[[26]int][]string{}

	for i := 0; i < len(strs); i++ {
		k := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			k[strs[i][j]-'a'] += 1
		}
		mp[k] = append(mp[k], strs[i])
	}

	resp := [][]string{}

	for _, v := range mp {
		resp = append(resp, v)
	}
	return resp
}

// @lc code=end
