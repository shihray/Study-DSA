/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
package leetcode

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
    m--
    n--
    for n >= 0{
        if m >= 0 && nums1[m] > nums2[n]{
            nums1[m + n + 1] = nums1[m]
            m--
        } else {
            nums1[m + n + 1] = nums2[n]
            n--
        }
    }
}
// @lc code=end

