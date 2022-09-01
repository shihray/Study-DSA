/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

package leetcode

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    ans := 0

	hash12Map := map[int]int{}

	for _, i1 := range nums1 {
		for _, i2 := range nums2 {
			hash12Map[i1+i2] += 1
		}
	}

	for _, i3 := range nums3 {
		for _, i4 := range nums4 {

			val := (i3 + i4) * -1

			if data, ok := hash12Map[val]; ok {
				ans += data
			}
		}
	}


	return ans
}
// @lc code=end

