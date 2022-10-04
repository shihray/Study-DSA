/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */
package q00658

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
    for len(arr) > k {
		if abs(arr[0] - x) > abs(arr[len(arr) - 1] - x) {
			arr = arr[1:]
		} else {
			arr = arr[:len(arr)-1]
		}
	}
	return arr
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
// @lc code=end

