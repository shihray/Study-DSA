/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    mergeArr := append(nums1, nums2...)
	sort.Ints(mergeArr)

	median := len(mergeArr) / 2
	if len(mergeArr)%2 == 0 {
		i, j := mergeArr[median], mergeArr[median-1]
		return float64((i + j)) / 2

	} else {
		return float64(mergeArr[median])
	}
}
// @lc code=end

