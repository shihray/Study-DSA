/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */
package q00540

// @lc code=start
func singleNonDuplicate(nums []int) int {
	dict := make(map[int]int)
	for _, num := range nums {
		dict[num]++
	}

	for v, d := range dict {
		if d == 1 {
			return v
		}
	}
	return 0
}

// @lc code=end
