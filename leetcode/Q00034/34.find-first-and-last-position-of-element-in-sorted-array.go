/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package q00034

// @lc code=start
func searchRange(nums []int, target int) []int {
    result := []int{-1,-1}
	
	if len(nums) == 0 {
		return result
	}

	return []int{findStart(nums, target), findEnd(nums, target)}
}

func findStart(nums []int, target int) int {
	start, end := 0 , len(nums)-1

	for start < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}


func findEnd(nums []int, target int) int {
	start, end := 0 , len(nums)-1

	for start < end {
		mid := (start + end) / 2 + 1
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid
		}
	}
	if nums[end] == target {
		return end
	}
	return -1
}
// @lc code=end

