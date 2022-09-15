/*
 * @lc app=leetcode id=1089 lang=golang
 *
 * [1089] Duplicate Zeros
 */
package q01089

// @lc code=start
func duplicateZeros(arr []int)  {
    zeros := 0

    for _, v := range arr {
        if v == 0 {
            zeros++
        }
    }

    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] == 0 {
            if zeros + i < len(arr) {
                arr[zeros + i] = 0
            }
            
            if zeros - 1 + i < len(arr) {
                arr[zeros - 1 + i] = 0
            }
            
            zeros--
        } else if i + zeros < len(arr) {
            arr[zeros + i] = arr[i]
        }
    }
}
// @lc code=end

