/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */
package leetcode

// @lc code=start
func findDisappearedNumbers(nums []int) []int {

	for _, v := range nums {
		
		offset := abs(v) - 1
		if nums[offset] > 0 {
			nums[offset] = nums[offset] * -1
		}
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			nums[count] = i+1
			count++
		}
	}
	return nums[:count]
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
// @lc code=end

