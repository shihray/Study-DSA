/*
 * @lc app=leetcode id=2256 lang=golang
 *
 * [2256] Minimum Average Difference
 */
package q02256

import "math"

// @lc code=start
func minimumAverageDifference(nums []int) int {
	sum1 := 0 
	for _, v := range nums {
		sum1 += v
	}

	idx, diff, sum2 := 0 ,float64(sum1), 0
	for i := 0; i < len(nums); i++ {
		sum2 += nums[i]
		sum1 -= nums[i]
		a := math.Floor(float64(sum2) / float64(i+1))
		var b float64
		if i != len(nums) - 1 {
			b = math.Floor(float64(sum1) / float64(len(nums) - i - 1))
		}
		currDiff := math.Abs(a - b)
		if currDiff < diff {
			diff = currDiff
			idx = i
		}
	}
	return idx
}
// @lc code=end

