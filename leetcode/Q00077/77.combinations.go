/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
package q00077

// @lc code=start
func combine(n int, k int) [][]int {
    
	resp := make([][]int, 0)

	var comb func(arr []int, start int) 
	comb = func(arr []int, start int) {
		if len(arr) == k {
			dst := make([]int, k)
			copy(dst, arr)
			resp = append(resp, dst)
			return
		}

		for i := start; i <= n; i++ {
			arr = append(arr, i)
			comb(arr, i + 1)
			arr = arr[:len(arr)-1]
		}
		return
	}

	comb(make([]int, 0), 1)
	return resp
}
// @lc code=end

