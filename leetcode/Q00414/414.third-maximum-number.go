/*
 * @lc app=leetcode id=414 lang=golang
 *
 * [414] Third Maximum Number
 */

package q00414

import "math"

// @lc code=start
func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	
	for _, v := range nums {
		if v == max1 || v == max2 || v == max3 {
			continue
		}
		switch{
		case v > max1:
			max2, max3, max1 = max1, max2, v
		case v > max2:
			max2, max3 = v, max2
		case v > max3:
			max3 = v
		}
	}

	if max3 == math.MinInt64 {
		return max1
	}
	return max3
}
// @lc code=end

