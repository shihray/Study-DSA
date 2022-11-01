/*
 * @lc app=leetcode id=1051 lang=golang
 *
 * [1051] Height Checker
 */
package q01051
const minHeight = 1
const maxHeight = 100

func heightCheckerSorting(heights []int) int {
    var m [maxHeight+1]int
    
    for _, h := range heights {
        m[h]++
    }
    
    k := 0 // index of heights
    count := 0
    
    for curr := minHeight; curr <= maxHeight; curr++ {
        for j := m[curr]; j > 0; j-- {
            if heights[k] != curr {
                count++
            }
            k++
        }    
    }
    
    return count
}
// @lc code=end

