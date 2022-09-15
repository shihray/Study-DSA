/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

package q00350

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
    minList, maxList := nums1, nums2
    checkMap := map[int]int{}
    var output []int
    if len(nums1) > len(nums2) {
        minList, maxList = nums2, nums1
    }
    for _, num := range minList {
        checkMap[num]++
    }
    for _, num := range maxList {
        if _, ok := checkMap[num]; ok {
            output = append(output, num)
            if checkMap[num] > 1 {
                checkMap[num] -= 1
            } else {
                delete(checkMap, num)
            }
        }
    }
    return output
}
// @lc code=end

