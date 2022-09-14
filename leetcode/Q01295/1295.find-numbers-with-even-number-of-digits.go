/*
 * @lc app=leetcode id=1295 lang=golang
 *
 * [1295] Find Numbers with Even Number of Digits
 */
package leetcode

import "strconv"

// @lc code=start
func findNumbers(nums []int) int {
    count := 0

	for i := 0; i < len(nums); i++ {
		if findEven(nums[i]) {
			count++
		}
	}

	return count
}

func findEven(num int) bool {
	
	str := strconv.Itoa(num)
	if len(str) % 2 == 0 {
		return true
	} 
	return false
}
// @lc code=end

