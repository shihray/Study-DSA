/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */
package q00075

// @lc code=start
func sortColors(nums []int) {

	size := len(nums)
	maxVal := 3

	count := make([]int, maxVal)

	for _, n := range nums {
		count[n] += 1
	}

	for i := 1; i < maxVal; i++ {
		count[i] += count[i-1]
	}

	origin := make([]int, size)
	copy(origin, nums)
	
	for i := size - 1; i >= 0; i -= 1 {
		n := origin[i]
 
		c := &count[n]
		*c -= 1
		nums[*c] = n
	}
}
// @lc code=end

