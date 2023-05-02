/*
 * @lc app=leetcode id=2336 lang=golang
 *
 * [2336] Smallest Number in Infinite Set
 */
package q2336

import "container/heap"

// @lc code=start
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}

type SmallestInfiniteSet struct {
	curr    int
	minHeap *MinHeap
	m       map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		curr:    1,
		minHeap: &MinHeap{},
		m:       make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.minHeap.Len() > 0 {
		v := heap.Pop(this.minHeap).(int)
		delete(this.m, v)
		return v
	}

	v := this.curr
	this.curr++
	return v
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.curr || this.m[num] {
		return
	}

	this.m[num] = true
	heap.Push(this.minHeap, num)
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
// @lc code=end
