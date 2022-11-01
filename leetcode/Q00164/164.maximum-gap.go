/*
 * @lc app=leetcode id=164 lang=golang
 *
 * [164] Maximum Gap
 */
package q00164

import "sort"

// @lc code=start
func maximumGap(nums []int) int {
	result := 0
	if len(nums) < 2 {
		return result
	}
	
	count := 0
    for i := 0 ; i < len(nums)-1 ; i++ {
        if nums[i] == nums[i+1] {
            count++
            if count == len(nums)-1{
                return 0
            }
        }
    }

	sort.Ints(nums)

	nums2 := unq(nums)
	sort.Ints(nums2)

	result = nums2[1] - nums2[0]
	
	for i := 2; i < len(nums2); i++ {
		g := nums[i] - nums[i - 1]
		if g > result {
			result = g
		}
	}
	return result
}

func unq(nums []int) []int {
	m := map[int]bool{}
	var res []int
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = true
			res = append(res, nums[i])
		}
	}
	return res
}

// @lc code=end
