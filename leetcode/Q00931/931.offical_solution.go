/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */
package q00931

import (
	"math"
)

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	resp := math.MaxInt32
	n := len(matrix)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp := minArr(matrix[i-1][maxArr(j-1,0)], matrix[i-1][j], matrix[i-1][minArr(j+1, n-1)])
			matrix[i][j] = matrix[i][j] + tmp
		}
	}

	for i := 0; i < n; i++ {
		resp = minArr(resp, matrix[n-1][i])
	}
    return resp
}

func minArr(vals ...int) int {
	resp := vals[0]
	for i := 1; i < len(vals); i++ {
		if resp > vals[i] {
			resp = vals[i]
		}
	}
	return resp
}

func maxArr(vals ...int) int {
    maxVal := math.MinInt32
    
    for _, val := range vals {
        if val > maxVal {
            maxVal = val
        }
    }
    
    return maxVal
}
// @lc code=end

