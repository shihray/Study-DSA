/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
    mList1 := map[string]int{}
	output := []string{}
	min := 2001  // max than two list total length
	for index, val := range list1 {
		mList1[val] = index
	} 
	
	for index, val := range list2 {
		if resp, ok := mList1[val]; ok {

			indexSum := resp + index

			if indexSum < min {
				min = indexSum
				output = []string{val}
			} else if indexSum == min {
				output = append(output, val)
			}
		}
	}
	return output
}
// @lc code=end

