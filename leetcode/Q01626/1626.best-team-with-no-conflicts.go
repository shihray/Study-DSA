/*
 * @lc app=leetcode id=1626 lang=golang
 *
 * [1626] Best Team With No Conflicts
 */
package q01626

import (
	"sort"
)

// @lc code=start
func bestTeamScore(scores []int, ages []int) int {
	l := len(scores)
	players := [][]int{}

	for i := 0; i < l; i++ {
		players = append(players, []int{scores[i], ages[i]})
	}

	sort.Slice(players, func(i, j int) bool {

		if players[i][0] == players[j][0] {
			return players[i][1] < players[j][1]
		}
		return players[i][0] < players[j][0]
	})

	maxAge := 0
	for i := 0; i < l; i++ {
		maxAge = max(maxAge, ages[i])
	}

	dp := make([]int, maxAge)
	resp := 0

	for _, player := range players {
		score := player[0]
		age := player[1]

		m := 0
		i := age

		for i > 0 {
			m = max(m, dp[i-1])
			i -= i & (-i)
		}

		dp[age-1] = m + score

		i = age
		for i <= len(dp) {
			dp[i-1] = max(dp[age-1], dp[i-1])
			i += i & (-i)
		}
		resp = max(resp, dp[age-1])
	}
	return resp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
