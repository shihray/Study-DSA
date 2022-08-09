/*
 * @lc app=leetcode id=1346 lang=golang
 *
 * [1346] Check If N and Its Double Exist
 */
package leetcode

// @lc code=start
func checkIfExist(arr []int) bool {

	existMap := map[int]bool{}

    for _, v := range arr {
		if existMap[v*2] || ( existMap[v/2] && v %2 == 0) {
			return true
		}
		existMap[v] = true
	}
	return false
}
// @lc code=end

func checkIfExist_hashMap(arr []int) bool {

	var hash = make(map[int]bool)
	var countZero int

	for _, val := range arr {
		hash[val] = true
		if val == 0 {
			countZero++
		}
	}

	for val := range hash {
		if hash[2*val] && val != 0 {
			return true
		}

		if val == 0 && countZero > 1 {
			return true
		}
	}
	return false
}