/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

package q00622

// @lc code=start
type MyCircularQueue struct {
    slice []int
	front int
	rear int
	size int
}


func Constructor(k int) MyCircularQueue {
    arr := make([]int, k)
	return MyCircularQueue{
		slice: arr,
		front: 0,
		rear: -1,
		size: k,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
		return false
	}
	this.rear++
	this.slice[this.rear % this.size] = value
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	
	if this.IsEmpty() {
		return false
	}
	this.front++
	return true
}


func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
		return -1
	}
    return this.slice[this.front % this.size]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
    return this.slice[this.rear % this.size]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.rear < this.front
}


func (this *MyCircularQueue) IsFull() bool {
	if this.rear - this.front == this.size-1 {
		return true
	}
	return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end