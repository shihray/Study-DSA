/*
 * @lc app=leetcode id=916 lang=golang
 *
 * [916] Word Subsets
 */
package leetcode

// @lc code=start
func wordSubsets(words1 []string, words2 []string) []string {
   
	counter, resp := map[rune]int{}, []string{}

	for _, word := range words2 {
		item := getRuneMap(word)
		for k, v := range item {
			counter[k] = maxVal(counter[k], v)
		}
	}

	for _, word := range words1 {
		tmp := getRuneMap(word)
		pass := true

		for k, v := range counter {
			if tmp[k] < v {
				pass = false
				break
			}
		}
		if pass {
			resp = append(resp, word)
		}
	}
	return resp 
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
} 

func getRuneMap(s string) map[rune]int {
	resp := map[rune]int{}

	for _, v := range s {
		resp[v]++
	}
	return resp
}
// @lc code=end

