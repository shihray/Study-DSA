/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */
package q01207

// @lc code=start
func uniqueOccurrences(arr []int) bool {
    intArr := map[int]int{}
    for _, i := range arr {
        intArr[i]++
    }
    
    intMap := map[int]int{}
    for _, v := range intArr {
        intMap[v]++
		if intMap[v] != 1 {
			return false
		}
    }
    return true
}
// @lc code=end

