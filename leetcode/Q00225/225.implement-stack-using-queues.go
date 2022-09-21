/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */
package q00225

// @lc code=start
type MyStack struct {
    stack []int
}


func Constructor() MyStack {
    return MyStack{}
}


func (this *MyStack) Push(x int)  {
    this.stack = append(this.stack, x)
}


func (this *MyStack) Pop() int {
	var result int
    if this.Empty() {
		return result
	}
	result = this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return result
}


func (this *MyStack) Top() int {
    var result int
    if this.Empty() {
		return result
	}
	return this.stack[len(this.stack)-1]
}


func (this *MyStack) Empty() bool {
    if len(this.stack) == 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

