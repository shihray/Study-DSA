/*
 * @lc app=leetcode id=2244 lang=golang
 *
 * [2244] Minimum Rounds to Complete All Tasks
 */
package q02244

// @lc code=start
func minimumRounds(tasks []int) int {
	count := map[int]int{}
	res := 0

	for _, v := range tasks {
		count[v]++
	}

	for _, v := range count {
		if v == 1 {
			return -1
		}

		if v%3 == 0 {
			res += v / 3
		} else {
			res += (1 + v/3)
		}
	}
	return res
}

// @lc code=end
