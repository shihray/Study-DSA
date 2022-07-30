/*
 * @lc app=leetcode id=890 lang=golang
 *
 * [890] Find and Replace Pattern
 */

// @lc code=start
func findAndReplacePattern(words []string, pattern string) []string {
    resp := []string{}

	for i := range words {
		if len(words[i]) != len(pattern) {
			continue
		}

		mapCurWord := make(map[byte]byte, len(words[i]))
		mapPattern := make(map[byte]byte, len(pattern))
		isValid := true

		for j := 0; j < len(words[i]); j++ {
			if _, ok := mapPattern[pattern[j]]; !ok {
                mapPattern[pattern[j]] = words[i][j]
            }
            if _, ok := mapCurWord[words[i][j]]; !ok {
                mapCurWord[words[i][j]] = pattern[j]
            }
            if mapCurWord[words[i][j]] != pattern[j] || mapPattern[pattern[j]] != words[i][j]{
                isValid = false
                break
            }
		}
		if isValid {
            resp = append(resp, words[i])
        }
	}
	return resp
}
// @lc code=end

