/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

package q00232

// @lc code=start
type MyQueue struct {
    queue []int
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.queue = append(this.queue, x)
}


func (this *MyQueue) Pop() int {
	var result int
    if this.Empty() {
		return result
	}
	result = this.queue[0]
	this.queue = this.queue[1:]
	return result
}


func (this *MyQueue) Peek() int {
    var result int
    if this.Empty() {
		return result
	}
	return this.queue[0]
}


func (this *MyQueue) Empty() bool {
    if len(this.queue) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

