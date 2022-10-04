/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */
package q00167

// @lc code=start
func twoSum(numbers []int, target int) []int {
	
	m := map[int]int{}
	
	for i := 0; i < len(numbers); i++ {
		m[numbers[i]] = i + 1
	}

	for i := 0; i < len(numbers); i++ {

		sub := target - numbers[i]
		if j, ok := m[sub]; ok {
			return []int{i + 1, j}
		}
	}
	return nil
}
// @lc code=end

