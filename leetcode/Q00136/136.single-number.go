/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	m := map[int]int{}
    
    for _, val := range nums {
        m[val]++
    }

	for index, val := range m {
		if val == 1 {
			return index 
		}
	}
	return 0
}
// @lc code=end

