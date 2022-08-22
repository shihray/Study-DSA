/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
    m := map[int]bool{}
	var resp []int

	for _, v := range nums1 {
		m[v] = true
	}

	for _, v := range nums2 {
		if item, ok := m[v]; ok && item {
			resp = append(resp, v)
			m[v] = false
		}
	}
	return resp
}
// @lc code=end

