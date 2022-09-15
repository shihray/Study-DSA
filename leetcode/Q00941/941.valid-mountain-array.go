/*
 * @lc app=leetcode id=941 lang=golang
 *
 * [941] Valid Mountain Array
 */
package q00941

// @lc code=start
func validMountainArray(arr []int) bool {
    
	front, back := 0, len(arr)

	if back < 3 {
		return false
	}

	for front + 1 < back && arr[front] < arr[front + 1] {
		front++
	}

	if front == 0 || front == (back - 1) {
		return false
	}

	for front+1 < back && arr[front] > arr[front+1]{
        front++
    }

	return front == back - 1 
}
// @lc code=end

