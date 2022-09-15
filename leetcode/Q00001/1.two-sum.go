/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package q00001

// @lc code=start
func twoSum(nums []int, target int) []int {

	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		in := target - nums[i]
		if index, ok := m[in]; !ok {
			m[nums[i]] = i
		} else {
			return []int{i, index}
		}
	}
	return nil
}
// @lc code=end

// func twoSum(nums []int, target int) []int {
	
// 	for i := 0; i < len(nums); i++ {
// 		output := []int{i}
// 		for j := i+1; j < len(nums); j++ {
// 			if nums[i] + nums[j] == target {
// 				output = append(output, j)
// 				return output
// 			}
// 		}
// 	}
// 	return nil
// }

