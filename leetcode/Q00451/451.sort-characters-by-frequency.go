/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */
package q00451

import (
	"sort"
)

// @lc code=start
func frequencySort(s string) string {
    countMap := make([]int, 255)

	for i := range s {
		countMap[s[i]]++
	}

	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if countMap[b[i]] == countMap[b[j]] {
			return b[i] > b[j]
		}
		return countMap[b[i]] > countMap[b[j]]
	})
	return string(b)
}
// @lc code=end

