/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */
package q00733

// @lc code=start
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    hit := image[sr][sc]
	if hit == color {
		return image
	}

	Flood(&image, sr, sc, color, hit)
	
	return image
}

func Flood(image *[][]int, sr, sc, color int, hit int) {

	m, n := len(*image), len((*image)[0])

	if sr < 0 || sr >= m || sc < 0 || sc >= n {
		return 
	}

	if (*image)[sr][sc] == hit {
		(*image)[sr][sc] = color

		Flood(image, sr + 1, sc, color, hit)
		Flood(image, sr - 1, sc, color, hit)
		Flood(image, sr, sc + 1, color, hit)
		Flood(image, sr, sc - 1, color, hit)
	}
}

// @lc code=end

