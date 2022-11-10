/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
package q00119

// @lc code=start
func getRow(rowIndex int) []int {
    var cur []int

	for i := 0; i <= rowIndex; i++ {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		cur = append(cur, 1)
		for j := 1; j < len(cur) - 1; j++ {
			cur[j] = tmp[j-1] + tmp[j]
		}
	}
	return cur
}
// @lc code=end

