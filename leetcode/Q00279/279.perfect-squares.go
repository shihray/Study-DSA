/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */
package q00279

// @lc code=start
func numSquares(n int) int {
    squares := []int{}

	for i := 1; i * i <= n; i++ {
		squares = append(squares, i*i)
	}

	queue := []int{n}
    visited := make(map[int]bool)
	breadthSize := 0
	
	for len(queue) > 0 {
		breadthSize++
		curLen := len(queue)

		for _, popped := range queue {
			for _, square := range squares {
				if popped == square {
					return breadthSize
				}

				remainder := popped - square
				if remainder > 0 && !visited[remainder] {
					visited[remainder] = true
					queue = append(queue, remainder)
				}
			} 
		}   
		queue = queue[curLen:]  
	}
	return -1
}
// @lc code=end

