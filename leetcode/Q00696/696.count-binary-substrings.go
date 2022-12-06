/*
 * @lc app=leetcode id=696 lang=golang
 *
 * [696] Count Binary Substrings
 */
package q00696

// @lc code=start
func countBinarySubstrings(s string) int {
	count := 0
    contiguous0, contiguous1 := 0, 0
    last := '2'
    for _, v := range s {
        if v == '0' {
            if v != last {
                contiguous0 = 0
            }
            contiguous0++
            if contiguous0 <= contiguous1 {
                count++
            }
        } else {
            if v != last {
                contiguous1 = 0
            }
            contiguous1++
            if contiguous1 <= contiguous0 {
                count++
            }
        }
        last = v
    }
    return count
}
// @lc code=end