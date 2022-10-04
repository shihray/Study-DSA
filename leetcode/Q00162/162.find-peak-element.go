/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */
package q00162

// @lc code=start
func findPeakElement(nums []int) int {
    start, end := 0, len(nums) - 1

	if len(nums) == 1 {
		return 0
	}

	for {
		mid := start + (end - start) / 2

		if mid == 0 {
			if nums[mid] > nums[mid + 1] {
				return mid
			}
			start = mid + 1
			continue
		}

		if mid == len(nums) - 1 {
			if nums[mid] > nums[mid - 1] {
				return mid
			}
			end = mid - 1
			continue
		}
		
		if nums[mid] > nums[mid -1] && nums[mid] > nums[mid + 1] {
			return mid
		}
		if nums[mid] > nums[mid - 1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
}
// @lc code=end

