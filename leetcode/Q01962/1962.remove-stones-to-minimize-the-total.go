/*
 * @lc app=leetcode id=1962 lang=golang
 *
 * [1962] Remove Stones to Minimize the Total
 */
package q01962

import (
	"container/heap"
	"math"
)

// @lc code=start

func minStoneSum(piles []int, k int) int {
	h := IntHeap(piles)
	heap.Init(&h)

	for i := 0; i < k; i++ {
		h[0] -= (h[0] / 2)
		heap.Fix(&h, 0)
	}

	sum := 0
	for h.Len() > 0 {
		sum += h.Pop().(int)
	}
	return sum
}

func floor(i int) int {
	return int(math.Floor(float64(i) / 2))
}

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h IntHeap) Len() int {
	return len(h)
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	tmp := old[len(*h)-1]
	*h = old[0 : len(*h)-1]
	return tmp
}

// @lc code=end
