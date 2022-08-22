/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */
package leetcode

// @lc code=start
func isHappy(n int) bool {
    m := make(map[int]bool)
    tmp := n
    for tmp != 1 {
        if _, ok := m[tmp]; ok {
            return false
        }
        m[tmp] = true
        tmp = helper(tmp)
    }
    return true
}

func helper(n int) int {
    rst := 0
    for n != 0 {
        rst += (n % 10) * (n % 10)
        n /= 10
    }
    return rst
}
// @lc code=end

