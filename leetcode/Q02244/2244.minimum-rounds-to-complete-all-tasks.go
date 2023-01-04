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
		for v > 0 {
			switch {
			case v%3 == 0:
				res, v = res+v/3, 0
			case v > 4:
				res++
				v -= 3
			case v%2 == 0:
				res, v = res+v/2, 0
			case v == 1:
				return -1
			}
		}
	}
	return res
}

// @lc code=end
