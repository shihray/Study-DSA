/*
 * @lc app=leetcode id=2306 lang=golang
 *
 * [2306] Naming a Company
 */
package q02306

// @lc code=start
func distinctNames(ideas []string) int64 {
	initialGroup := map[string][26]bool{}
	count := [26]int{}
	same := [26][26]int{}
	val := 0
	s := ""
	result := int64(0)

	for _, idea := range ideas {
		s = string(idea[1:])
		val = int(idea[0] - 'a')
		count[val]++
		if _, ok := initialGroup[s]; !ok {
			initialGroup[s] = [26]bool{}
		}
		chars := initialGroup[s]
		chars[val] = true
		initialGroup[s] = chars
	}

	for _, chars := range initialGroup {
		for i := range chars {
			if !chars[i] {
				continue
			}
			for j := i + 1; j < 26; j++ {
				if !chars[j] {
					continue
				}
				same[i][j]++
			}
		}
	}
	for i := range count {
		if count[i] == 0 {
			continue
		}
		for j := i + 1; j < 26; j++ {
			if count[j] == 0 {
				continue
			}
			val = (count[i] - same[i][j]) * (count[j] - same[i][j])
			result += int64(val)
		}
	}

	return result * 2
}

// @lc code=end
