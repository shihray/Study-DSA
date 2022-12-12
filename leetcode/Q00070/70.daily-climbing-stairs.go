/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package q00070

// @lc code=start
func climbStairsV2(n int) int {
    m := map[int]int{
		0:1,
		1:1,
	}
	return helper(n, m)
}

func helper(n int, m map[int]int) int {
	if n <= 1 {
		return m[n]
	} else if v, exist := m[n]; exist {
		return v
	} else {
		m[n] = helper(n - 1, m) + helper(n - 2, m)
		return m[n]
	}
}

// @lc code=end

