/*
 * @lc app=leetcode id=2225 lang=golang
 *
 * [2225] Find Players With Zero or One Losses
 */
package q02225

import "math"

// @lc code=start
func findWinners(matches [][]int) [][]int {
    var losserCount = make([]int, 100001)

	for i := 0; i < 100001; i++ {
		losserCount[i] = math.MaxInt
	}

	for i := 0; i < len(matches); i++ {
		if losserCount[matches[i][0]] == math.MaxInt {
			losserCount[matches[i][0]] = 0
		}
		if losserCount[matches[i][1]] == math.MaxInt {
			losserCount[matches[i][1]] = 1
		} else {
			losserCount[matches[i][1]]++
		}
	}

	resp := make([][]int, 2)
	resp[0], resp[1] = make([]int, 0), make([]int, 0)
	for k, v := range losserCount {
		if v == 0 {
			resp[0] = append(resp[0], k)
		}
		if v == 1 {
			resp[1] = append(resp[1], k)
		}
	}
	return resp
}
// @lc code=end

