/*
 * @lc app=leetcode id=1299 lang=golang
 *
 * [1299] Replace Elements with Greatest Element on Right Side
 */
package leetcode

// @lc code=start
func replaceElements(arr []int) []int {
    
	maxVal := -1

	for i := len(arr)-1; i >= 0; i-- {	
		currData := arr[i]	
		arr[i] = maxVal
		maxVal = max(currData, maxVal)
	}
	return arr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// @lc code=end

