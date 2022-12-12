/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */
package q00347

// @lc code=start
func topKFrequentV2(nums []int, k int) []int {
    
	m := map[int]int{}
	bucket := make([][]int, len(nums)+1)

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if bucket[v] == nil {
			bucket[v] = make([]int, 0)
		}
		bucket[v] = append(bucket[v], k)
	}

	res := make([]int, k)
	cnt := 0
	for i := len(bucket) - 1; i >= 0 && cnt < k; i-- {
		if bucket[i] != nil {
			for _, i := range bucket[i] {
				res[cnt] = i
                cnt ++
			}
		}
	}
	return res
}
// @lc code=end

