/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */
package q00953

// @lc code=start
func isAlienSorted(words []string, order string) bool {

	orderMap := []int{}

	for i, v := range order {
		orderMap[v] = i
	}

	var isOrder func(str1, str2 string) bool

	isOrder = func(str1, str2 string) bool {
		minlen := min(str1, str2)
		for i := 0; i < minlen; i++ {
			if str1[i] == str2[i] {
				continue
			}

			if orderMap[rune(str1[i])] > orderMap[rune(str2[i])] {
				return false
			} else {
				return true
			}
		}
		return len(str1) <= len(str2)
	}

	for i := 1; i < len(words); i++ {
		if !isOrder(words[i-1], words[i]) {
			return false
		}
	}
	return true
}

func min(a, b string) int {
	if len(a) < len(b) {
		return len(a)
	}
	return len(b)
}

// @lc code=end
