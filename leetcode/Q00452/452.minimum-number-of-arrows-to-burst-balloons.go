/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */
package q00452

import "sort"

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			count++
			end = points[i][1]
		}
	}
	return count
}

// @lc code=end
