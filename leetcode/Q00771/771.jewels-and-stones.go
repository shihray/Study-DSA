/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 */
package q00771

// @lc code=start
func numJewelsInStones(jewels string, stones string) int {
    m := map[byte]int{}
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = 0
	}
	count := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			count++
		}
	}
	return count
}
// @lc code=end

