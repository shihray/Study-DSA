/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package q00070

// @lc code=start
func climbStairs(n int) int {
	memo := map[int]int{
		0: 1,
		1: 1,
	}
	return climb(n, memo)
}

func climb(i int, memo map[int]int) int {

	if i <= 1 {
		return memo[i]
	} else if val, exist := memo[i]; exist {
		return val
	} else {
		memo[i] = climb(i-1, memo) + climb(i-2, memo)
		return memo[i]
	}
}

// @lc code=end
