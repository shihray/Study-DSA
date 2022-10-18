/*
 * @lc app=leetcode id=1200 lang=golang
 *
 * [1200] Minimum Absolute Difference
 */
package q01200

import "sort"

// @lc code=start
func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)

	mixAbs := arr[1] - arr[0]

	for i := 1; i < len(arr) - 1; i++ {
		tmp := arr[i + 1] - arr[i]
		if tmp < mixAbs {
			mixAbs = tmp
		} 
	}

	resp := [][]int{}
	for i := 0; i < len(arr) - 1; i++ {
		tmp := arr[i + 1] - arr[i]
		if tmp == mixAbs {
			resp = append(resp, []int{arr[i], arr[i + 1]})
		} 
	}
	return resp
}
// @lc code=end

